[![Go Reference](https://pkg.go.dev/badge/github.com/jreisinger/gokatas.svg)](https://pkg.go.dev/github.com/jreisinger/gokatas)
[![Go Report Card](https://goreportcard.com/badge/github.com/jreisinger/gokatas)](https://goreportcard.com/report/github.com/jreisinger/gokatas)

> Practice yourself, for heaven's sake, in little things; and thence proceed to greater. -- Epictetus (Discourses IV.i)

# Go katas

Katas (形) are practiced in martial arts as a way to memorize and perfect the
movements being executed. Let's try the same with code. The approach is pretty
low-tech. Go katas is basically a
[list](https://pkg.go.dev/github.com/jreisinger/gokatas#section-directories) of
packages and commands that you should be rewriting from scratch or partially.
There's a command to visualize your progress:

```
$ go run cmd/katas.go
Kata           Last done  Done
----           ---------  ----
areader      11 days ago    1x
bytecounter   8 days ago    1x
clock2        8 days ago    1x
----                      ----
3                            3
```

## Why

I've been learning to program in Go. I work in the sysadmin/devops and security
areas so I usually don't get to program every day. But I still want to keep my
skills fresh. Maybe even improve them. I use gokatas as one of the ways to
achieve this.

## How

It's important to practice regularly, to create a habit. Start by setting a
goal that you can meet, e.g. 15 minutes every day before work. At first it's
fine even if you only read through one of the katas. After some time it will
require much less will power and you will become more familiar with the code.

Try to write (and understand!) and then run or test the code before checking
the solution (by running `git diff`).

If you feel comfortable enough with a kata stop practicing it (for some time)
and pick up one that interests you and is slightly beyond your current ability.

## Initial setup

1) [Install Go](https://go.dev/doc/install).

2) [Fork](https://github.com/jreisinger/gokatas/fork) and then clone the repo: `git clone git@github.com:<you>/gokatas.git`.

3) Start practicing:

```
cd gokatas
> katas.md # if you are not me :-)
go doc
```
