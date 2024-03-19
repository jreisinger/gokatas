[![Go Report Card](https://goreportcard.com/badge/github.com/jreisinger/gokatas)](https://goreportcard.com/report/github.com/jreisinger/gokatas)
[![test and scan](https://github.com/jreisinger/gokatas/actions/workflows/test-scan.yaml/badge.svg)](https://github.com/jreisinger/gokatas/actions/workflows/test-scan.yaml)

# Go katas

Katas (形) are practiced in martial arts as a way to internalize and perfect the
techniques so they can be executed and adapted under different circumstances,
without thought or hesitation. Let's try the same with Go code.

The approach is pretty low-tech. Go katas is a list of directories containing
brief, well written Go programs. Select one of them and try to understand it.
Then be rewriting it partially or from scratch until you feel comfortable with it.
Use `git diff` to see how you are doing. There's a command to show the katas 
with some info and your progress:

```
$ go run ./cmd/gokatas -c 2
Kata              Last done    Done  Level     Topics
----              ---------    ----  -----     ------
boring/boring     0 days ago     1x  beginner  concurrency, design
boring/channel    0 days ago     1x  beginner  goroutines, channels
areader           3 days ago     2x  beginner  interfaces, io.Reader
----                           ----
3                                4x 
```

Initial setup:

1) [Install Go](https://go.dev/doc/install).

2) [Fork](https://github.com/jreisinger/gokatas/fork) this repo so you have your own copy.
  
3) Clone the fork to your computer: `git clone git@github.com:<you>/gokatas.git`.

4) Clear contents of `katas.md` (unless you are me :-) and start practicing.

It's important to practice regularly because repetition creates habits, and
habits are what enable mastery. Start by taking baby steps. Set a goal that you
can meet, e.g. 10 minutes every day before work. At first it's fine even if you
only read through one of the beginner level katas. Use documentation, a search
engine or an AI model if you don't understand something. After some time it will
require much less will power to practice. Your programming moves will start
looking simpler and smoother. If you feel comfortable enough with a kata, stop
practicing it and pick another one that interests you and is slightly beyond
your current ability.
