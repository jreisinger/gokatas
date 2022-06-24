[![Go Reference](https://pkg.go.dev/badge/github.com/jreisinger/gokatas.svg)](https://pkg.go.dev/github.com/jreisinger/gokatas)
[![Go Report Card](https://goreportcard.com/badge/github.com/jreisinger/gokatas)](https://goreportcard.com/report/github.com/jreisinger/gokatas)

> Practice yourself, for heaven's sake, in little things; and thence proceed to greater. -- Epictetus

# Go katas

Katas (形) are practiced in martial arts as a way to memorize and perfect the movements being executed. Let's try the same with code. The approach is pretty low-tech. Go katas is basically a [list](https://pkg.go.dev/github.com/jreisinger/gokatas#section-directories) of packages and commands that you should be rewriting from scratch or at least partially. There's also a command to visualize your progress:

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

Go katas should be useful for folks trying to learn to program in Go and/or for those who don't get to program every day but still want to keep their skills fresh (e.g. sysadmins, devops or security people).

It's important to practice regularly, to create a habit. Start by taking baby steps, set a daily minimum (e.g. 20 minutes). After some time it will require much less will power and it will be harder to stop.

## Initial setup

1) [Install Go](https://go.dev/doc/install).

2) [Fork](https://github.com/jreisinger/gokatas/fork) and then clone the repo: `git clone git@github.com:<you>/gokatas.git`.

3) Start practicing:

```
cd gokatas
> katas.md # if you are not me :-)
go doc
```
