// Sample functions using switch statement.
//
// Level: beginner
// Topics: switch, flow control
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println(sign(1))
	greet()
	system()
	saturday()
	fmt.Println(sqlQuote(true))
}

// sign says whether the number is positive, negative or zero.
func sign(x int) string {
	switch { // tagless switch, equivalent to "switch true"
	case x > 0:
		return "positive"
	case x < 0:
		return "negative"
	default:
		return "zero"
	}
}

// greet gives a greeting based on the time of the day
// (https://tour.golang.org/flowcontrol/11).
func greet() {
	now := time.Now()
	switch {
	case now.Hour() < 12:
		fmt.Println("Good morning!")
	case now.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

// system prints what system Go is running on
// (https://tour.golang.org/flowcontrol/9).
// go tool dist list | cut -d'/' -f1 | sort | uniq
func system() {
	fmt.Print("Go is running on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Mac OS X.")
	case "linux":
		fmt.Println("Linux.")
	case "openbsd":
		fmt.Println("OpenBSD.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

// saturday prints when's Saturday
// (https://tour.golang.org/flowcontrol/10).
func saturday() {
	fmt.Print("Saturday is ")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today:
		fmt.Println("today!")
	case today + 1:
		fmt.Println("tomorrow.")
	default:
		fmt.Println("too far away...")
	}
}

// sqlQuote shows a type switch (GoPL ch 7.13).
func sqlQuote(x interface{}) string {
	switch x := x.(type) {
	case nil: // x == nil
		return "NULL"
	case int, uint:
		return fmt.Sprintf("%d", x) // x has type interface{} here.
	case bool:
		if x {
			return "TRUE"
		}
		return "FALSE"
	case string:
		return sqlQuoteString(x)
	default:
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}
}

func sqlQuoteString(x string) string {
	return "NOT IMPLEMENTED"
}
