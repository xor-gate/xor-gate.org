---
date: "2019-10-06"
tags: ["protocols", "embedded", "messagepack", "protobuf", "nanopb", "msgpack"]
title: "Embedded Systems using JSON, MessagePack or Protocol Buffer?"
draft: true
---

The [MessagePack](http://msgpack.org) binary protocol runs also on constrained systems which is a better choice for embedded systems then JSON. Because it is deterministic as it includes an header per item. It is by design exchangable with JSON, only it is a binary protocol. I added [C11 Generic](https://en.cppreference.com/w/c/language/generic) support into the [github.com/ludocode/msgpack](https://github.com/ludocode/msgpack) which makes the code more readable and reliable. It also is natural less bytes on the wire than JSON because of a simple compression trick. And like [Protocol Buffers](https://en.wikipedia.org/wiki/Protocol_Buffers) it doesn't need a schema to make use of the data. Which means Protocol Buffers is even less bytes on the wire. There is even the [nanopb](https://github.com/nanopb/nanopb) embedded implementation available but it is more work to setup and maintain then a plain protocol codec because it also generates code.
