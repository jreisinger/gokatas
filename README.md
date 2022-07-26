[![Go Reference](https://pkg.go.dev/badge/github.com/jreisinger/gokatas.svg)](https://pkg.go.dev/github.com/jreisinger/gokatas)
[![Go Report Card](https://goreportcard.com/badge/github.com/jreisinger/gokatas)](https://goreportcard.com/report/github.com/jreisinger/gokatas)

> Practice yourself, for heaven's sake, in little things; and thence proceed to greater. -- Epictetus (Discourses IV.i)

# Go katas

Katas (形) are practiced in martial arts as a way to memorize and perfect the movements being executed. Let's try the same with code. The approach is pretty low-tech. Go katas is basically a [list](https://pkg.go.dev/github.com/jreisinger/gokatas#section-directories) of packages and commands that you should be rewriting from scratch or partially. There's a command to visualize your progress:

```
$ go run cmd/katas.go
Kata                         Last done       Count
----                         ---------       -----
areader              11 days ago (Fri)           1
bytecounter           8 days ago (Mon)           1
clock2                8 days ago (Mon)           1
----                                         -----
3                                                3
```

## Why

Go katas should be useful for folks trying to learn to program in Go and/or for those who don't get to program every day but still want to keep their skills fresh (e.g. sysadmins, devops or security people). Go katas can also be used as a copy/paste reference.

## How

It's important to practice regularly. So you need to create a habit. Start by
setting a goal that you can meet, e.g. 15 minutes every day before work. At
first it's fine even if you only read through one of the katas. After some time
it will require much less will power, it will become natural.

If you feel comfortable enough with a kata stop practicing it for some time and pick up one that is slighly beyond your current ability. Try to write, run or test the code and (most probably) debug it before looking up the solution.

## Initial setup

1) [Install Go](https://go.dev/doc/install).

2) [Fork](https://github.com/jreisinger/gokatas/fork) and then clone the repo: `git clone git@github.com:<you>/gokatas.git`.

3) Start practicing:

```
cd gokatas
> katas.md # if you are not me :-)
go doc
```
