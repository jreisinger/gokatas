// Sortbooks sorts a collection of books. To sort a collection of elements you
// need to define a type for this collection. This type needs to have the
// methods that satisfy the sort.Interface interface type.
//
// It's not necessary to use a pointer to book. However when swapping many
// elements it might be faster this way since pointer is always only a machine
// word in size (usually 32 or 64 bits).
//
// 	b := book{"The Go Programming Language", "Kernighan, Donovan", 2015}
// 	p := &b
// 	format := "type: %10T, size: %2d bytes\n"
// 	fmt.Printf(format, b, unsafe.Sizeof(b)) // type:  main.book, size: 40 bytes
// 	fmt.Printf(format, p, unsafe.Sizeof(p)) // type: *main.book, size:  8 bytes
//
// Based on github.com/adonovan/gopl.io/blob/master/ch7/sorting/.
package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
)

type book struct {
	Title   string
	Authors authors
	Year    int // when published
}

type authors []string

func (a authors) String() string {
	return strings.Join(a, ", ")
}

var books = []*book{
	{"The Lord of the Rings", authors{"Tolkien"}, 1954},
	{"The Phoenix Project", authors{"Kim", "Behr", "Spafford"}, 2013},
	{"The Go Programming Language", authors{"Kernighan", "Donovan"}, 2015},
}

type ByYear []*book

func (x ByYear) Len() int           { return len(x) }
func (x ByYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x ByYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func main() {
	sort.Sort(sort.Reverse(ByYear(books)))
	printBooks(books)
}

func printBooks(books []*book) {
	const format = "%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Authors", "Year")
	fmt.Fprintf(tw, format, "-----", "-------", "----")
	for _, b := range books {
		// you don't have to derefence here like (*b).title
		fmt.Fprintf(tw, format, b.Title, b.Authors, b.Year)
	}
	tw.Flush() // calculate column widths and print table
}
