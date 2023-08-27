---
date: "2023-08-27"
title: "The digital afterlife"
---

![the-digital-afterlife-image](/images/the-digital-afterlife/dan-meyers-f1WMJR8pLqo-unsplash.jpg)

[- image by Dan Meyers - Unsplash.com](https://unsplash.com/photos/f1WMJR8pLqo)

What about the digital traces we leave behind on the internet after we pass away? When your most important accounts are secured with [multi-factor authentication](https://en.wikipedia.org/wiki/Multi-factor_authentication) it is impossible to recover. People who are left behind probably have no clue how to access your most important internet accounts and data of value like family pictures.

My most important account is my private Google account. Where I have stored my contacts and e-mail, and a backup of data of value for me. The backup is synchronized every night from my FreeBSD server with [`rclone`](https://rclone.org/). Because I only use exclusivly Apple products for desktop, laptop and smartphone purposes I also have a secured Apple ID.

Apple has the ["Legacy Contact"](https://support.apple.com/en-us/HT212360). And Google has the ["Inactive Account Manager"](https://myaccount.google.com/inactive). So your account and data can be recovered with your death certificate or is released after account inactivity.

I'm a little paranoid so for every internet account I create a different password and store it in a Keepass encrypted database. On macOS I use MacPass. On Windows at work I use [KeePass XC](https://keepassxc.org/) with [Chrome browser plugin](https://keepassxc.org/docs/KeePassXC_GettingStarted#_setup_browser_integration).

I also backup the QR-codes used with 2-Factor Authentication smartphone apps like Google Authenticator or Microsoft Authentication. And extract the TOTP-secret from the QR-code to place in Keepass so I can generate the 6-number security codes on my iMac and MacBook with [MacPass](https://macpassapp.org/). When my Smartphone gets lost or broken, I can still generate the "military-grade" security codes. But luckily Google Authenticator now has a cloud backup feature (which I don't use).

I'm to paranoid for storing my passwords and secrets in the cloud like [Lastpass announced a security incident December 2022](https://blog.lastpass.com/2022/12/notice-of-recent-security-incident/).

For extracting the TOTP-secret data from stored QR-code image/screenshot on a mac. Assuming you are a power user and installed the [Homebrew package manager](https://brew.sh/). No you don't want to use an online service to extract the TOTP-secret from the QR-code for obvious reasons!

```
jerry@Jerrys-iMac ~ % brew install zbar
jerry@Jerrys-iMac Downloads % zbarimg google-qr-2fa.png
QR-Code:otpauth://totp/Google%3AXXX%40gmail.com?secret=ecbazdoxjob5b56vhpjdmz6eioeqousc&issuer=Google
scanned 1 barcode symbols from 1 images in 0,03 seconds
```

As you can see the [QR-code](https://en.wikipedia.org/wiki/QR_code) contains a special formatted URL. Where most important `key=value` is `secret=XXX`. This secret must be inserted in you [MacPass/KeePass XC OTP secret field](https://keepassxc.org/docs/KeePassXC_UserGuide#_adding_totp_to_an_entry). This `otpauth` protocol URL specification is [not well standarized](https://shkspr.mobi/blog/2022/05/why-is-there-no-formal-specification-for-otpauth-urls/). See also [`github.com/google/google-authenticator` wiki Key Uri Format](https://github.com/google/google-authenticator/wiki/Key-Uri-Format)

For the security researchers among us, making a physical copy of the TOTP-secret can also be a security hole. But I encrypt the QR-code images in my Keepass database. Or print them out on paper. When losing it, you can lose your account! A potential security problem is far worse than out-locking yourself from your most important accounts. Choose your "security poison".

**The just-in-case box**

You should prepare for the worst and create a small box with printed information like:

* Internet (e-mail) accounts like: Google, Apple, Microsoft
* Most important phone numbers like work, clubs, friends, family etc.
* List of subscriptions with client or contract numbers
* List of contacts to send a death notice to
* Smartphone pincode
* Smartphone SIM-card pin/puk code
* Computer password(s)
* For the digital (crypto) paranoid:
  * [Yubi hardware encryption key](https://www.yubico.com/)
  * Server and/or desktop full disk encryption keys
  * Screenshots or copies of 2-factor QR-codes ([TOTP secrets](https://en.wikipedia.org/wiki/Time-based_one-time_password))
  * SSH private keys
* ... and the list goes on and on

The box could contain an USB stick FAT32 formatted (synced every 31 march on [World Backup Day](https://www.worldbackupday.com/en)) for maximum OS compatibility. My USB stick also contains an encrypted Keepass password and account database.

So how are you prepared for the digital afterlife?
