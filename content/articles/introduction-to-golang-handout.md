---
title: "Introduction to Golang Handout"
date: 2019-10-06T18:47:22+02:00
tags: ["golang", "go"]
draft: true
---

This article is the handout of the [Introduction to Golang](/talks/introduction-to-golang.html) talk.

# History

Between 2007 and 2012

- Project starts at Google in 2007 (by Griesemer, Pike, Thompson)
- Open source release in November 2009 ([[https://www.youtube.com/watch?v=rKnDgT73v8s][youtube]]/[[https://talks.golang.org/2009/go_talk-20091030.pdf][slides]])
- More than 250 contributors join the project
- Version 1.0 release in March 2012

2018 and beyond

- [[https://blog.golang.org/go2-here-we-come][Go 2 announced]] in November 2018
- [[https://blog.golang.org/go1.13][Version 1.13]] release in September 2019 
- Solved the dependency management problem as [[https://blog.golang.org/using-go-modules][go modules]]

# The Go approach to Concurrency

- Go provides a way to write systems and servers
as concurrent, garbage-collected processes
(goroutines) with support from the language and
run-time
- Language takes care of goroutine management,
memory management
- Growing stacks, multiplexing of goroutines onto
threads is done automatically
- Concurrency is hard without garbage collection
- Garbage collection is hard without the right
language

* The Go approach to Dependencies

Construction speed depends on managing
dependencies. Explicit dependencies in source allow:
- fast compilation
- fast linking

The Go compiler pulls transitive dependency type
info from the object file - but only what it needs.

If A.go depends on B.go depends on C.go:
- compile C.go, B.go, then A.go.
- to compile A.go, compiler reads B.o not C.o.

At scale, this can be a huge speedup.

* Golang for Embedded Systems?

Go is not designed for CPU & memory constrained systems.

- [[https://tinygo.org][TinyGo]]
- [[https://github.com/ziutek/emgo][emgo]]
- [[https://gobot.io][Gobot]]
- [[https://embd.kidoman.io][Golang embedded framework]]

* Quicklinks

- [[https://try.golang.org][try.golang.org]]
- [[https://blog.golang.org][blog.golang.org]]
- [[https://github.com/avelino/awesome-go][github.com/avelino/awesome-go]]
- [[https://www.golang-book.com/books/intro][An Introduction to Programming in Go]]
- [[https://youtube.com/c/justforfunc][JustForFunc: Programming in Go]]
