---
date: "2019-04-02"
tags: ["protocols", "embedded", "stm32", "stlink", "arm"]
title: "ST-Link AP != 0 access by TG"
---

I received a very nice private mail from TG which fidled around with the STlink and the ARM Access Port (AP) beyond the AP 0.
With the normal stlink API it is not possible to access other access ports, and TG found out how you can use it.

> From: TG </br>
> Subject: ST-Link AP!=0 access </br>
> Date: April 2, 2019 at 1:10:29 AM MDT </br>
> 
> Hello Texane and Jerry,
> 
> I hope you'll pardon this random email!
> 
> I'm a firmware developer working on some python scripts to drive the ST-Link debugger on my STM32H743 Nucleo board.  I used OpenOCD and the texane/stlink project to understand the protocol and send some USB commands via pyusb and it works quite well.  However, I was disappointed to see that there appeared to be no ability to access the AP registers nor to perform memory accesses to AP1, 2, 3, etc. using the ST-Link probe.
> 
> While fiddling with the 16-byte zero-padded command buffers, I discovered that the memory and debug register read/write commands actually take an APSEL parameter at the end of the command (which everyone normally zero-pads so everyone accesses AP0).  Intrigued by this, I enabled a USB sniffer and watched what some of the ST tools do and was also able to also find commands that perform read operations on AP registers.  On a wild guess I looked at the next undocumented command number and was able to find a command that performs a write operation on an AP register.
> 
> I thought you might find this of interest since it allows the ST-Link probe to do some lower-level accesses (and probably manually set up 16-bit accesses via the AP TAR/DRW/BDn registers which seem impossible through the probe's high-level API which only seems to support 8- or 32-bit accesses).
> 
> The commands are as follows (all multi-byte values are little-endian and the firmware version on my probe is V2J29M18):
> 
> Read AP Reg </br>
> 0xF2 0x45 0xnnnn 0xyy </br>
> nnnn - APSEL value (probably only allows 0-255) </br>
> yy - register offset in AP space (00=CSW, 04=TAR, etc.) - may actually be a 16-bit value but it is zero-padded at the end of the command </br>
> The returned value is at offset 4 in the response, much like Read Debug 32. </br>
> 
> Write AP Reg </br>
> 0xF2 0x46 0xnnnn 0xyyyy 0xvvvvvvvv </br>
> nnnn - APSEL value </br>
> yyyy - register offset in AP space </br>
> vvvvvvvv - value to write to the AP register </br>
> 
> Read Mem 32 - same as normal with APSEL tacked on at the end </br>
> 0xF2 0x07 0xyyyyyyyy 0xzzzz 0xnn </br>
> nn - APSEL value (may also be 16-bit but again is zero-padded) </br>
> yyyyyyyy - memory address </br>
> zzzz - number of words to read </br>
> 
> Write Mem 32 - same as normal with APSEL tacked on at the end </br>
> 0xF2 0x08 0xyyyyyyyy 0xzzzz 0xnn </br>
> nn - APSEL value </br>
> 
> Read Debug 32 </br>
> 0xF2 0x36 0xyyyyyyyy 0xnn </br>
> nn - APSEL value </br>
> 
> Write Debug 32 </br>
> 0xF2 0x35 0xyyyyyyyy 0xvvvvvvvv 0xnn </br>
> nn - APSEL value </br>
> 
> On the STM32H743 there is an APB-AP on APSEL 2 which I suppose may be present on other STM32's as well.  Through this AP you can access the DBGMCU at its documented address 0xE00E1000.  However, if you use the standard ST-Link memory access commands they will clobber AP2's CSW.DbSwEnable bit.  This then prevents accesses to the DBGMCU via AP0 (and presumably the CPU) at 0x5C001000.  Access can be restored by manually setting AP2's CSW.DbSwEnable with the new Write AP Reg command.
> 
> I especially like this because it allows one to read th AP ROM Table BASE address at offset 0xF8 in the AP - making it possible to scan the ROM tables using the ST-Link debugger without having to manually probe hard-coded device addresses.  To properly scan some ROM tables you will also need to enable the debug clocks via `DBGMCU_CR` - with a smart ROM table probing algorithm you can do this as you descend the hierarchy and identify components.
> 
> I hope you find this useful and thank you for your open source work which has been very helpful to learn how the protocol works.
> 
> TG
