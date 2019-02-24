---
date: "2019-02-24"
tags: ["protocols", "embedded", "redis", "resp"]
title: "The RESP3 protocol"
---

I'm a long time fan of the [Redis](http://redis.io) project. During some deep diving into the internals while developing a microservice using Redis I figured out the simple and elegant binary safe ASCII protocol named [RESP](https://redis.io/topics/protocol) (REdis Serialization Protocol). It is so easy i would love to run it on a memory and CPU contrained embedded system. So i started to implement it in the C99 programming language. The only problem with constrained systems is you don't have "infinite" amount of storage and memory theirfor it is contrained. So we need to be carefull how to implement it in a nice way. 

[antirez](https://github.com/antirez/resp3) has written the specification and is available on [github](https://github.com/antirez/resp3). He also wrote a [blog post](http://antirez.com/news/125) on the subject. Still there are some concerns about the Redis project and forward progress which also includes the new backward compatible RESP3 protocol. The full post can be [read here](http://antirez.com/news/126).

I will release the embedded respv3 library opensource on Github in the near future under the MIT License. Stay tuned!
