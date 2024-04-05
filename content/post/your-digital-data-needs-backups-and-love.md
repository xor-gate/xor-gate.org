---
title: Your digital data needs backups and love
date: 2024-03-30
tags:
  - BSD
  - open-source
  - FreeBSD
draft: false
---
![harddisk-closeup](/images/your-digital-data-needs-backups-and-love/patrick-lindenberg-1iVKwElWrPA-unsplash.jpg)
Photo by [Patrick Lindenberg](https://unsplash.com/@heapdump?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash) on [Unsplash](https://unsplash.com/photos/photo-of-optical-disc-drive-1iVKwElWrPA?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash)

Back in the days when I started using a computer Microsoft Windows 95 was new and shiny. Cloud, smartphones and always connected to the internet was not a thing. We needed to backup our work on floppy disks. Many years later the USB memory stick was available to make our data portable and could store many many more data then fitted on a single floppy. I can still remember buying a 128 MByte Kingston USB memory stick at a computer fair. It was expensive, like in 40 dutch guldens or so. During that time (re)writable compact disks I also did use for data which only hold around 700Mbyte. And later on the DVD which could hold around 4 GByte per disk. I had a illegal copy of Microsoft Encarta encyclopaedia on CD and used it so much and with the low quality disk I had I literally used the disk until it died. I only had just one copy, so it was lost. Luckily it was not own generated content/data/media, so it was no big loss. 

Many years later my computer had a windows crash it did probably run Windows 2000 (NT) or Window 98 i don't remember exactly. I did reboot my system and it started to "eat" my data on the second personal data harddrive. Some sort of disk check/recovery which "restored" the state of the filesystem. You can almost guess it, my data was gone. Only some disk "check and repair" folder/files where left. At that time I didn't know about [digital forensics like PhotoRec](https://en.wikipedia.org/wiki/PhotoRec), so it was lost forever.

When I started to backup some own data it was stored on DVDs  and used them for a year of 3-5. And didn't had a duplicate on other type of media somewhere. Then it broke and was unreadable. Luckily life didn't depend on the personal data, but if it was used for a company or serious personal projects it could also be a serious financial or intellectual property loss.

Nowadays I'm a little paranoid about the safety of my personal data. I try to automated all the things so I can sleep and the computer is at work for me to safeguard the data. Yes offline (and offsite) backups need manual labor, but it is worth the effort. 

Tomorrow 31 march 2024 it is World Backup Day. In case somebody reads this, when you have something of value as digital data always have multiple copies. When media dies or something bad happens you will thank yourself for putting effort in a (simple) backup strategy.

Bad things can happen for example when a [Crypto Locker malware](https://en.wikipedia.org/wiki/CryptoLocker) just changes files on a shared (corporate) network drive and your files will be taken hostage. Even if you pay the (major) randsom you don't know if you get the decryption keys and the tool/algorithm to reveal your "secured" data again. It is even worse when the randsom payment has a time limit, then the bad guys will publish and/or sell your data on the dark-web. If your ID card was on that network share you have serious problem (again). Restore from a crypto locker could be "as easy as" putting a filesystem snapshot back. Which can be done in a zippy with ZFS and BTRFS copy-on-write filesystems. But if you have a windows server and the samba share, you need another solution to restore the historical data. It even gets more worse when a network drive is like 30TB or something, when loading the data from an offsite backup (if you have that in place). You will have downtime!

Digital data needs backups and love. Even if you read this after World Backup Day. It is food for thought.

As a bonus a short explanation of my personal data setup and software I use:

* Dedicated NAS hosted at home
	* ["Vanilla" FreeBSD OS](https://en.wikipedia.org/wiki/FreeBSD) with [ZFS filesystem](https://en.wikipedia.org/wiki/ZFS)
	* [Syncthing](https://en.wikipedia.org/wiki/Syncthing) "central server"
	* [rclone](https://en.wikipedia.org/wiki/Rclone) nightly cronjob for offsite backup sync to Google Drive
	* 14 TB ZFS RAID-Z mirror data integrity and automatic snapshots
* iMac, macbook desktops: Uses Syncthing to my NAS
* iPhone smartphone: [Google drive](https://en.wikipedia.org/wiki/Google_Drive) for file storage "read-only" as data is nightly synced from the NAS

***"Keep them bytes safe!"***