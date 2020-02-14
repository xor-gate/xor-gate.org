---
title: "FreeBSD Jails"
date: 2020-02-13T22:03:25+01:00
tags: ["freebsd", "docker", "containers"]
---

Because I switched [this year](/post/switching-from-linux-to-freebsd) from [Linux](https://en.wikipedia.org/wiki/Linux) to [FreeBSD](https://en.wikipedia.org/wiki/FreeBSD) I don't have Linux Containers (e.g Docker). A
 much simpler containerization is FreeBSD [jail(8)](https://www.freebsd.org/cgi/man.cgi?query=jail) which has been around
since the FreeBSD 4.x which is released over [20 years](https://en.wikipedia.org/wiki/FreeBSD_version_history#FreeBSD_4) ago!
A note must be made It is not as advanced as the Linux counterpart with its [complex namespaces](https://medium.com/@teddyking/linux-namespaces-850489d3ccf) implemented as [CLONE_*](https://github.com/torvalds/linux/blob/9065e0636036e4f8a6f65f9c34ed384e4b776273/tools/include/uapi/linux/sched.h#L28-L33) system calls to limit memory, CPU resources etc, etc. This post is based on FreeBSD release 12.1 so your results may vary.

## Starting, stopping and restarting a single jail

Control of individual jails uses the system `service` command with the `jails` service and its parameters.

```
service jail start <jailname>
service jail stop <jailname>
service jail restart <jailname>
```

## Showing all running jails

For showing all running jails with its IP, Hostname and Path the `jls` command is used.

```
$ jls
root@pineapple:~/dl # jls
   JID  IP Address      Hostname                      Path
     8  192.168.1.101   dev.volkstuin.club            /var/jails/cgadmin-dev
     9  192.168.1.102   ci.xor-gate.org               /var/jails/ci
```

## Enter a jail shell

To enter a jail with an interactive shell the `jexec` command is used.

```
root@pineapple:~/dl # jexec ci
root@ci:/ # uname -a
FreeBSD ci.xor-gate.org 12.1-RELEASE FreeBSD 12.1-RELEASE r354233 GENERIC  amd64
```

## Fixing DNS not working in the jail

A vanilla FreeBSD tarball run under a jail when using a jail configured static IP doesn't
have the DNS resolver configured. FreeBSD DNS resolver uses the `/etc/resolv.conf` file to
get the DNS server list to make the actual request. So when using `drill` to get a DNS A record
IP it gives an error.

```
root@dev:/ # drill xor-gate.org
Warning: Could not create a resolver structure: Could not open the files ((null))
Try drill @localhost if you have a resolver running on your machine.
```

You must create the `/etc/resolv.conf` file with e.g `nameserver 8.8.8.8` pointing to
 the primary Google DNS server.
