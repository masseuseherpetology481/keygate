package handler

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/tabloy/keygate/internal/config"
	"github.com/tabloy/keygate/internal/middleware"
	"github.com/tabloy/keygate/internal/model"
	"github.com/tabloy/keygate/internal/service"
	"github.com/tabloy/keygate/internal/store"
	"github.com/tabloy/keygate/pkg/response"
)

// setSecureCookie sets a cookie with SameSite=Lax for CSRF protection.
func setSecureCookie(c *gin.Context, name, value string, maxAge int, path string, secure, httpOnly bool) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     name,
		Value:    value,
		MaxAge:   maxAge,
		Path:     path,
		Secure:   secure,
		HttpOnly: httpOnly,
		SameSite: http.SameSiteLaxMode,
	})
}

type AuthHandler struct {
	Store  *store.Store
	Config *config.Config
	Email  *service.EmailService
}

func (h *AuthHandler) Me(c *gin.Context) {
	uid, _ := c.Get("user_id")
	uidStr, ok := uid.(string)
	if !ok || uidStr == "" {
		response.Unauthorized(c, "unauthorized")
		return
	}
	user, err := h.Store.FindUserByID(c, uidStr)
	if err != nil {
		response.NotFound(c, "user not found")
		return
	}
	response.OK(c, gin.H{
		"id": user.ID, "email": user.Email, "name": user.Name,
		"avatar_url": user.AvatarURL, "is_admin": user.IsAdmin(), "role": user.Role,
	})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	userID, _ := c.Get("user_id")
	if raw, err := c.Cookie("refresh_token"); err == nil && raw != "" {
		h.Store.DeleteRefreshToken(c, hashToken(raw))
	}
	// Delete user's other refresh tokens to fully invalidate session
	if uid, ok := userID.(string); ok && uid != "" {
		h.Store.DeleteUserRefreshTokens(c, uid)
		h.Store.Audit(c, &model.AuditLog{
			Entity: "session", EntityID: uid, Action: "logout",
			ActorType: "user", ActorID: uid, IPAddress: c.ClientIP(),
		})
	}
	setSecureCookie(c, "session", "", -1, "/", h.Config.IsProduction(), true)
	setSecureCookie(c, "refresh_token", "", -1, "/api/v1/auth/refresh", h.Config.IsProduction(), true)
	response.OK(c, gin.H{"status": "logged_out"})
}

func (h *AuthHandler) Refresh(c *gin.Context) {
	raw, err := c.Cookie("refresh_token")
	if err != nil || raw == "" {
		response.Unauthorized(c, "no refresh token")
		return
	}

	tokenHash := hashToken(raw)
	rt, err := h.Store.FindRefreshToken(c, tokenHash)
	if err != nil {
		response.Unauthorized(c, "invalid refresh token")
		return
	}

	user, err := h.Store.FindUserByID(c, rt.UserID)
	if err != nil {
		response.Unauthorized(c, "user not found")
		return
	}

	// Rotate: delete old, issue new
	h.Store.DeleteRefreshToken(c, tokenHash)
	h.issueSession(c, user)
	response.OK(c, gin.H{"status": "refreshed"})
}

func (h *AuthHandler) issueSession(c *gin.Context, user *model.User) {
	// JWT includes admin claim for convenience, but the authoritative check
	// happens at request time via DB role lookup in SessionAuth middleware.
	token, _ := middleware.IssueJWT(
		h.Config.JWTSecret, user.ID, user.Email, user.Name,
		user.IsAdmin(), 24*time.Hour,
	)
	setSecureCookie(c, "session", token, 24*3600, "/", h.Config.IsProduction(), true)

	// Long-lived refresh token (30 days)
	rawRefresh := randomHex(32)
	refreshHash := hashToken(rawRefresh)
	expiresAt := time.Now().Add(30 * 24 * time.Hour)
	_ = h.Store.CreateRefreshToken(c, user.ID, refreshHash, expiresAt)
	setSecureCookie(c, "refresh_token", rawRefresh, 30*24*3600, "/api/v1/auth/refresh", h.Config.IsProduction(), true)
}

func hashToken(raw string) string {
	h := sha256.Sum256([]byte(raw))
	return hex.EncodeToString(h[:])
}

func (h *AuthHandler) Providers(c *gin.Context) {
	devLogin := h.Config.IsDevLoginAllowed()
	response.OK(c, gin.H{"dev_login": devLogin, "otp": true})
}

// DevLogin is a development-only endpoint that creates a session without email OTP.
// Security: This endpoint is ONLY available when BOTH conditions are met:
//  1. ENVIRONMENT is explicitly set to "development"
//  2. The server is NOT listening on a public interface (checked via BASE_URL)
//
// This prevents accidental exposure in production deployments.
func (h *AuthHandler) DevLogin(c *gin.Context) {
	if !h.Config.IsDevLoginAllowed() {
		response.NotFound(c, "not found")
		return
	}
	// Block dev-login on public-facing hosts
	base := h.Config.BaseURL
	if !strings.Contains(base, "localhost") && !strings.Contains(base, "127.0.0.1") {
		response.NotFound(c, "not found")
		return
	}

	var req struct {
		Email string `json:"email" binding:"required"`
		Name  string `json:"name"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "email is required")
		return
	}
	if req.Name == "" {
		req.Name = "Dev User"
	}

	user := &model.User{Email: req.Email, Name: req.Name}
	if err := h.Store.UpsertUser(c, user); err != nil {
		response.Internal(c)
		return
	}
	user, err := h.Store.FindUserByEmail(c, req.Email)
	if err != nil {
		response.Internal(c)
		return
	}

	h.issueSession(c, user)
	h.Store.Audit(c, &model.AuditLog{
		Entity: "session", EntityID: user.ID, Action: "login",
		ActorType: "dev_login", ActorID: user.ID, IPAddress: c.ClientIP(),
		Changes: map[string]any{"email": user.Email},
	})
	// Auto-promote if email is in ADMIN_EMAILS and user is currently just 'user'
	if h.Config.IsAdminEmail(user.Email) && user.Role == model.RoleUser {
		_ = h.Store.SetUserRole(c, user.ID, model.RoleAdmin)
		user.Role = model.RoleAdmin
	}

	response.OK(c, gin.H{
		"status": "ok", "email": user.Email, "name": user.Name,
		"is_admin": user.IsAdmin(), "role": user.Role,
	})
}

func randomHex(n int) string {
	b := make([]byte, n)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
