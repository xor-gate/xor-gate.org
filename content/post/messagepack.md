---
date: "2019-02-24"
tags: ["protocols", "embedded", "messagepack"]
title: "Embedded Systems using JSON or MessagePack?"
draft: true
---

The [MessagePack](http://msgpack.org) protocol runs also on constrained systems which is a better choice for embedded systems then JSON. Because it is deterministic as it includes an header per item. It is by design exchangable with JSON, only it is a binary protocol. I added [C11 Generic](https://en.cppreference.com/w/c/language/generic) support into the [github.com/ludocode/msgpack](https://github.com/ludocode/msgpack) which makes the code more readable and reliable.
