---
title: "Computers and Machines"
date: 2022-01-09T11:03:51+01:00
draft: true
---

My hosts are named after [tropical fruits](https://en.wikipedia.org/wiki/Category:Tropical_fruit)

{{< toc >}}

## coconut.xor-gate.org

### Specifications

* Apple iMac (Retina 4K, 21.5-inch, Late 2015) (Model ID: iMac16,2)
* CPU 3,1 GHz Quad-Core Intel Core i5
* Memory 16 GB 1867 MHz DDR3
* macOS Monterey (v12)

## pineapple.xor-gate.org

Self build low-power file/download server

### Specifications

* Mainboard Supermicro [A1SAM-2550F](https://www.supermicro.com/products/motherboard/Atom/X10/A1SAM-2550F.cfm)
* 16GB DDR3 non-ECC (4x 4GB)
* 128GB SSD for the FreeBSD operating system
* 2TB WD Red HDD for data
* 4x Gigabit LAN
* 1x Management ethernet interface (Supermicro IPMI iKVM)

### TODO ZFS pool

* <https://wintelguy.com/zfs-calc.pl>
* <https://blog.programster.org/zfs-create-disk-pools>

### Console output

**CPU information**

```
root@pineapple:/home/jerry # sysctl -a hw.model
hw.model: Intel(R) Atom(TM) CPU  C2550  @ 2.41GHz
```

## banana.xor-gate.org

### Specifications

* MSI Wind Box DE200 nettop custom form-factor (MS-7467)
* 2TB mobile HDD
* 2GB DDR2 memory

### Console output

**CPU information**

```
root@banana:/home/jerry # sysctl -a hw.model
hw.model: Intel(R) Atom(TM) CPU  330   @ 1.60GHz
```

**PCI information**

```
root@banana:/home/jerry # pciconf -lv
hostb0@pci0:0:0:0:	class=0x060000 rev=0x02 hdr=0x00 vendor=0x8086 device=0x2770 subvendor=0x1462 subdevice=0x7467
    vendor     = 'Intel Corporation'
    device     = '82945G/GZ/P/PL Memory Controller Hub'
    class      = bridge
    subclass   = HOST-PCI
hdac0@pci0:0:27:0:	class=0x040300 rev=0x01 hdr=0x00 vendor=0x8086 device=0x27d8 subvendor=0x1462 subdevice=0x7467
    vendor     = 'Intel Corporation'
    device     = 'NM10/ICH7 Family High Definition Audio Controller'
    class      = multimedia
    subclass   = HDA
pcib1@pci0:0:28:0:	class=0x060400 rev=0x01 hdr=0x01 vendor=0x8086 device=0x27d0 subvendor=0x1462 subdevice=0x7467
    vendor     = 'Intel Corporation'
    device     = 'NM10/ICH7 Family PCI Express Port 1'
    class      = bridge
    subclass   = PCI-PCI
pcib2@pci0:0:28:1:	class=0x060400 rev=0x01 hdr=0x01 vendor=0x8086 device=0x27d2 subvendor=0x1462 subdevice=0x7467
    vendor     = 'Intel Corporation'
    device     = 'NM10/ICH7 Family PCI Express Port 2'
    class      = bridge
    subclass   = PCI-PCI
pcib3@pci0:0:28:2:	class=0x060400 rev=0x01 hdr=0x01 vendor=0x8086 device=0x27d4 subvendor=0x1462 subdevice=0x7467
    vendor     = 'Intel Corporation'
    device     = 'NM10/ICH7 Family PCI Express Port 3'
    class      = bridge
    subclass   = PCI-PCI
uhci0@pci0:0:29:0:	class=0x0c0300 rev=0x01 hdr=0x00 vendor=0x8086 device=0x27c8 subvendor=0x1462 subdevice=0x7467
    vendor     = 'Intel Corporation'
    device     = 'NM10/ICH7 Family USB UHCI Controller'
    class      = serial bus
    subclass   = USB
uhci1@pci0:0:29:1:	class=0x0c0300 rev=0x01 hdr=0x00 vendor=0x8086 device=0x27c9 subvendor=0x1462 subdevice=0x7467
    vendor     = 'Intel Corporation'
    device     = 'NM10/ICH7 Family USB UHCI Controller'
    class      = serial bus
    subclass   = USB
uhci2@pci0:0:29:2:	class=0x0c0300 rev=0x01 hdr=0x00 vendor=0x8086 device=0x27ca subvendor=0x1462 subdevice=0x7467
    vendor     = 'Intel Corporation'
    device     = 'NM10/ICH7 Family USB UHCI Controller'
    class      = serial bus
    subclass   = USB
uhci3@pci0:0:29:3:	class=0x0c0300 rev=0x01 hdr=0x00 vendor=0x8086 device=0x27cb subvendor=0x1462 subdevice=0x7467
    vendor     = 'Intel Corporation'
    device     = 'NM10/ICH7 Family USB UHCI Controller'
    class      = serial bus
    subclass   = USB
ehci0@pci0:0:29:7:	class=0x0c0320 rev=0x01 hdr=0x00 vendor=0x8086 device=0x27cc subvendor=0x1462 subdevice=0x7467
    vendor     = 'Intel Corporation'
    device     = 'NM10/ICH7 Family USB2 EHCI Controller'
    class      = serial bus
    subclass   = USB
pcib4@pci0:0:30:0:	class=0x060401 rev=0xe1 hdr=0x01 vendor=0x8086 device=0x244e subvendor=0x1462 subdevice=0x7467
    vendor     = 'Intel Corporation'
    device     = '82801 PCI Bridge'
    class      = bridge
    subclass   = PCI-PCI
isab0@pci0:0:31:0:	class=0x060100 rev=0x01 hdr=0x00 vendor=0x8086 device=0x27b8 subvendor=0x1462 subdevice=0x7467
    vendor     = 'Intel Corporation'
    device     = '82801GB/GR (ICH7 Family) LPC Interface Bridge'
    class      = bridge
    subclass   = PCI-ISA
atapci0@pci0:0:31:1:	class=0x01018a rev=0x01 hdr=0x00 vendor=0x8086 device=0x27df subvendor=0x1462 subdevice=0x7467
    vendor     = 'Intel Corporation'
    device     = '82801G (ICH7 Family) IDE Controller'
    class      = mass storage
    subclass   = ATA
atapci1@pci0:0:31:2:	class=0x01018f rev=0x01 hdr=0x00 vendor=0x8086 device=0x27c0 subvendor=0x1462 subdevice=0x7467
    vendor     = 'Intel Corporation'
    device     = 'NM10/ICH7 Family SATA Controller [IDE mode]'
    class      = mass storage
    subclass   = ATA
ichsmb0@pci0:0:31:3:	class=0x0c0500 rev=0x01 hdr=0x00 vendor=0x8086 device=0x27da subvendor=0x1462 subdevice=0x7467
    vendor     = 'Intel Corporation'
    device     = 'NM10/ICH7 Family SMBus Controller'
    class      = serial bus
    subclass   = SMBus
re0@pci0:1:0:0:	class=0x020000 rev=0x03 hdr=0x00 vendor=0x10ec device=0x8168 subvendor=0x1462 subdevice=0x7467
    vendor     = 'Realtek Semiconductor Co., Ltd.'
    device     = 'RTL8111/8168/8411 PCI Express Gigabit Ethernet Controller'
    class      = network
    subclass   = ethernet
none0@pci0:2:0:0:	class=0x028000 rev=0x10 hdr=0x00 vendor=0x10ec device=0x8172 subvendor=0x1462 subdevice=0x6897
    vendor     = 'Realtek Semiconductor Co., Ltd.'
    device     = 'RTL8191SEvB Wireless LAN Controller'
    class      = network
vgapci0@pci0:3:0:0:	class=0x030000 rev=0x00 hdr=0x00 vendor=0x1002 device=0x9552 subvendor=0x1002 subdevice=0x9552
    vendor     = 'Advanced Micro Devices, Inc. [AMD/ATI]'
    device     = 'RV710/M92 [Mobility Radeon HD 4330/4350/4550]'
    class      = display
    subclass   = VGA
hdac1@pci0:3:0:1:	class=0x040300 rev=0x00 hdr=0x00 vendor=0x1002 device=0xaa38 subvendor=0x1002 subdevice=0xaa38
    vendor     = 'Advanced Micro Devices, Inc. [AMD/ATI]'
    device     = 'RV710/730 HDMI Audio [Radeon HD 4000 series]'
    class      = multimedia
    subclass   = HDA
```

## mango.xor-gate.org

[Vorke V1](http://www.vorke.com/project/vorke-v1-2/)

### Specifications

* Custom form-factor Vorke V1 mainboard
* 4GB DDR3L
* 64GB mSata SSD for FreeBSD operating system (model FORESEE 64GB SSD)
* 512 GB intel SSD for data (model INTEL SSDSC2KW512G8)
* Gigabit LAN (PCI-e)
* 2x USB3
* 2x USB2

### Console output

**CPU information**

```
root@mango:/home/jerry # sysctl -a hw.model
hw.model: Intel(R) Celeron(R) CPU  J3160  @ 1.60GHz
```

**PCI information**

```
root@mango:/home/jerry # pciconf -lv
hostb0@pci0:0:0:0:	class=0x060000 rev=0x35 hdr=0x00 vendor=0x8086 device=0x2280 subvendor=0x8086 subdevice=0x7270
    vendor     = 'Intel Corporation'
    device     = 'Atom/Celeron/Pentium Processor x5-E8000/J3xxx/N3xxx Series SoC Transaction Register'
    class      = bridge
    subclass   = HOST-PCI
vgapci0@pci0:0:2:0:	class=0x030000 rev=0x35 hdr=0x00 vendor=0x8086 device=0x22b1 subvendor=0x8086 subdevice=0x7270
    vendor     = 'Intel Corporation'
    device     = 'Atom/Celeron/Pentium Processor x5-E8000/J3xxx/N3xxx Integrated Graphics Controller'
    class      = display
    subclass   = VGA
ahci0@pci0:0:19:0:	class=0x010601 rev=0x35 hdr=0x00 vendor=0x8086 device=0x22a3 subvendor=0x8086 subdevice=0x7270
    vendor     = 'Intel Corporation'
    device     = 'Atom/Celeron/Pentium Processor x5-E8000/J3xxx/N3xxx Series SATA Controller'
    class      = mass storage
    subclass   = SATA
xhci0@pci0:0:20:0:	class=0x0c0330 rev=0x35 hdr=0x00 vendor=0x8086 device=0x22b5 subvendor=0x8086 subdevice=0x7270
    vendor     = 'Intel Corporation'
    device     = 'Atom/Celeron/Pentium Processor x5-E8000/J3xxx/N3xxx Series USB xHCI Controller'
    class      = serial bus
    subclass   = USB
none0@pci0:0:26:0:	class=0x108000 rev=0x35 hdr=0x00 vendor=0x8086 device=0x2298 subvendor=0x8086 subdevice=0x7270
    vendor     = 'Intel Corporation'
    device     = 'Atom/Celeron/Pentium Processor x5-E8000/J3xxx/N3xxx Series Trusted Execution Engine'
    class      = encrypt/decrypt
hdac0@pci0:0:27:0:	class=0x040300 rev=0x35 hdr=0x00 vendor=0x8086 device=0x2284 subvendor=0x8086 subdevice=0x7270
    vendor     = 'Intel Corporation'
    device     = 'Atom/Celeron/Pentium Processor x5-E8000/J3xxx/N3xxx Series High Definition Audio Controller'
    class      = multimedia
    subclass   = HDA
pcib1@pci0:0:28:0:	class=0x060400 rev=0x35 hdr=0x01 vendor=0x8086 device=0x22c8 subvendor=0x8086 subdevice=0x7270
    vendor     = 'Intel Corporation'
    device     = 'Atom/Celeron/Pentium Processor x5-E8000/J3xxx/N3xxx Series PCI Express Port'
    class      = bridge
    subclass   = PCI-PCI
pcib2@pci0:0:28:1:	class=0x060400 rev=0x35 hdr=0x01 vendor=0x8086 device=0x22ca subvendor=0x8086 subdevice=0x7270
    vendor     = 'Intel Corporation'
    device     = 'Atom/Celeron/Pentium Processor x5-E8000/J3xxx/N3xxx Series PCI Express Port'
    class      = bridge
    subclass   = PCI-PCI
isab0@pci0:0:31:0:	class=0x060100 rev=0x35 hdr=0x00 vendor=0x8086 device=0x229c subvendor=0x8086 subdevice=0x7270
    vendor     = 'Intel Corporation'
    device     = 'Atom/Celeron/Pentium Processor x5-E8000/J3xxx/N3xxx Series PCU'
    class      = bridge
    subclass   = PCI-ISA
ichsmb0@pci0:0:31:3:	class=0x0c0500 rev=0x35 hdr=0x00 vendor=0x8086 device=0x2292 subvendor=0x8086 subdevice=0x7270
    vendor     = 'Intel Corporation'
    device     = 'Atom/Celeron/Pentium Processor x5-E8000/J3xxx/N3xxx SMBus Controller'
    class      = serial bus
    subclass   = SMBus
iwm0@pci0:1:0:0:	class=0x028000 rev=0x83 hdr=0x00 vendor=0x8086 device=0x08b3 subvendor=0x8086 subdevice=0x0070
    vendor     = 'Intel Corporation'
    device     = 'Wireless 3160'
    class      = network
re0@pci0:2:0:0:	class=0x020000 rev=0x0c hdr=0x00 vendor=0x10ec device=0x8168 subvendor=0x10ec subdevice=0x0123
    vendor     = 'Realtek Semiconductor Co., Ltd.'
    device     = 'RTL8111/8168/8411 PCI Express Gigabit Ethernet Controller'
    class      = network
    subclass   = ethernet
```

