---
title: "How I Started Learning Golang"
date: 2020-04-11T09:38:13+02:00
tags: ["go", "golang", "programming", "protocols", "ssh"]
---

Learning a new programming language is not a matter reading a book and you reach expert level. You need to get your hands dirty and
create something with it. As I'm a fond enthusiast of Open-Source software and active on [Github](https://github.com) I
started creating the [SSH File Transfer Protocol (sftp)](https://en.wikipedia.org/wiki/SSH_File_Transfer_Protocol) backend for Afero which is a golang package for file system
abstraction. My initial [Github pull request](https://github.com/spf13/afero/pull/48) dates back to 26 Dec 2015 so it is
some time ago. The project was stale for some time [and
now](https://github.com/spf13/afero/pull/157#ref-issue-126007661) it seems to be moving forward again which is nice to
see. I'm for a long time in protocols and especially interested in how you can layer applications and functionality on
top of [Secure Shell (SSH)](https://en.wikipedia.org/wiki/Secure_Shell) protocol.

It's also the reason why I started implementing the [SSHFP]() DNS record resolving based on RFCs in my own
[github.com/xor-gate/sshfp](https://github.com/xor-gate/sshfp) golang package. So Machine-to-Machine can verify out-of-band fingerprint of SSH servers
without any human interaction.

Be safe, use encryption!
