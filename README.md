[![Go Reference](https://pkg.go.dev/badge/github.com/jreisinger/gokatas.svg)](https://pkg.go.dev/github.com/jreisinger/gokatas)
[![Go Report Card](https://goreportcard.com/badge/github.com/jreisinger/gokatas)](https://goreportcard.com/report/github.com/jreisinger/gokatas)

# Go katas

Katas (形) are practiced in martial arts as a way to memorize and perfect the
movements being executed. Let's try the same with code. The approach is pretty
low-tech. Go katas is basically a
[list](https://pkg.go.dev/github.com/jreisinger/gokatas#section-directories) of
packages and commands that you should be rewriting from scratch or partially.
There's a command to show katas and your progress:

```
$ go run cmd/katas.go
Kata               Last done      Done  Level         Topics
----               ---------      ----  -----         ------
areader            0 days ago      51x  beginner      interfaces, io.Reader
noescape           0 days ago       6x  intermediate  html/template, constants, security
geometry           3 days ago      11x  intermediate  methods, math
<... snip ...>
----                              ----
48                                1340
```

## Why

I've been learning to program in Go. I work in the sysadmin/devops and security
areas so I usually don't get to program every day. But I still want to keep my
coding skills fresh. Maybe even improve them. I use gokatas as one of the ways
to achieve this.

## How

It's important to practice regularly, to create a habit. Start by setting a goal
that you can meet, e.g. 45 minutes every day before work. At first it's fine
even if you only read through one of the katas. After some time it will require
much less will power and you will become more familiar with the code.  If you
feel comfortable enough with a kata stop practicing it (for some time) and pick
up one that interests you and is slightly beyond your current ability.

## Initial setup

1) [Install Go](https://go.dev/doc/install).

2) [Fork](https://github.com/jreisinger/gokatas/fork) and then clone the repo: `git clone git@github.com:<you>/gokatas.git`.

3) Start practicing:

```
cd gokatas
go doc
```