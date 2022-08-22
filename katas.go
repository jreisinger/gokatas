package gokatas

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"text/tabwriter"
	"time"
)

// KatasFile is a MarkDown file to track katas you've done. It looks like this:
//
// 	* 2022-04-25: bytecounter, clock2
// 	* 2022-04-22: areader
const KatasFile = "katas.md"

// Kata represents a programming kata.
type Kata struct {
	Name       string
	Count      int
	LastDoneOn time.Time
}

// Get gets katas from the KatasFile.
func Get() ([]Kata, error) {
	f, err := os.Open(KatasFile)
	if err != nil {
		return nil, err
	}

	// Regexes
	kataLineRE := regexp.MustCompile(`^\s*\*\s*([0-9]{4}\-[0-9]{2}\-[0-9]{2}):\s*(.+)$`)
	comaRE := regexp.MustCompile(`\s*,\s*`)

	katas := make(map[string]Kata) // name to Kata

	s := bufio.NewScanner(f)
	for s.Scan() {
		lineParts := kataLineRE.FindStringSubmatch(s.Text())
		if lineParts == nil {
			continue
		}

		date, katasStr := lineParts[1], lineParts[2]
		doneOn, err := time.Parse("2006-01-02", date)
		if err != nil {
			return nil, err
		}

		for _, name := range comaRE.Split(katasStr, -1) {
			if name == "" {
				continue
			}
			if kata, ok := katas[name]; ok {
				kata.Count++
				if doneOn.After(kata.LastDoneOn) {
					kata.LastDoneOn = doneOn
				}
				katas[name] = kata
			} else {
				kata.Name = name
				kata.Count = 1
				kata.LastDoneOn = doneOn
				katas[name] = kata
			}
		}
	}
	if s.Err() != nil {
		return nil, s.Err()
	}

	var ks []Kata
	for name := range katas {
		ks = append(ks, katas[name])
	}

	return ks, nil
}

// Print prints table with statistics about katas. Only katas lastDoneDaysAgo or
// later are shown. Katas are sorted by when last done or by count.
func Print(katas []Kata, lastDoneDaysAgo int, sortByCount bool) {
	const format = "%v\t%v\t%5v\n"

	// Print header.
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Kata", "Last done", "Count")
	fmt.Fprintf(tw, format, "----", "---------", "-----")

	// Print lines.
	var katasCount int
	var totalCount int
	sortKatas(katas, &sortByCount)
	for _, k := range katas {
		if !show(k, lastDoneDaysAgo) {
			continue
		}

		katasCount++
		totalCount += k.Count

		fmt.Fprintf(tw, format, k.Name, formatLastDoneOn(k.LastDoneOn, sortByCount), k.Count)
	}
	// Print footer.
	fmt.Fprintf(tw, format, "----", "", "-----")
	fmt.Fprintf(tw, format, katasCount, "", totalCount)

	tw.Flush() // calculate column widths and print table
}

type customSort struct {
	katas []Kata
	less  func(x, y Kata) bool
}

func (x customSort) Len() int           { return len(x.katas) }
func (x customSort) Less(i, j int) bool { return x.less(x.katas[i], x.katas[j]) }
func (x customSort) Swap(i, j int)      { x.katas[i], x.katas[j] = x.katas[j], x.katas[i] }

// sortKatas first sorts by how recently the kata was done then by kata name.
func sortKatas(katas []Kata, countSort *bool) {
	sort.Sort(customSort{katas, func(x, y Kata) bool {
		if *countSort {
			if x.Count != y.Count {
				return x.Count > y.Count
			}
		} else {
			if x.LastDoneOn != y.LastDoneOn {
				return x.LastDoneOn.After(y.LastDoneOn)
			}
		}
		if x.Name != y.Name {
			return x.Name < y.Name
		}
		return false
	}})
}

// show decides when to show a kata. Negative lastDoneDaysAgo returns true.
func show(k Kata, lastDoneDaysAgo int) bool {
	if lastDoneDaysAgo < 0 {
		return true
	}
	t := time.Now().Add(-time.Hour * 24 * time.Duration(lastDoneDaysAgo+1))
	return k.LastDoneOn.After(t)
}

// formatLastDoneOn formats the time.
func formatLastDoneOn(lastDoneOn time.Time, sortByCount bool) string {
	daysAgo := int(time.Since(lastDoneOn).Hours() / 24)
	w := "day"
	if daysAgo != 1 {
		w += "s"
	}
	return fmt.Sprintf("%d %s ago", daysAgo, w)
}
