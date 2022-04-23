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

// PrintStats prints table with statistics about katas you've done.
func PrintStats(katasFile string, showAllKatas, sortByCount bool) error {
	var katas []kata
	katas, err := parseFile(katasFile)
	if err != nil {
		return err
	}
	stats := getStats(katas)
	printStats(stats, &showAllKatas, &sortByCount)
	return nil
}

// kata represents a programming kata.
type kata struct {
	name   string
	doneOn time.Time
}

// statistics represents statistics about a kata you've done.
type statistics struct {
	name       string
	count      int
	lastDoneOn time.Time
}

type customSort struct {
	s    []statistics
	less func(x, y statistics) bool
}

func (x customSort) Len() int           { return len(x.s) }
func (x customSort) Less(i, j int) bool { return x.less(x.s[i], x.s[j]) }
func (x customSort) Swap(i, j int)      { x.s[i], x.s[j] = x.s[j], x.s[i] }

// sortStats first sorts by how recently the kata was done then by kata name.
func sortStats(stats []statistics, countSort *bool) {
	sort.Sort(customSort{stats, func(x, y statistics) bool {
		if *countSort {
			if x.count != y.count {
				return x.count < y.count
			}
		} else {
			if x.lastDoneOn != y.lastDoneOn {
				return y.lastDoneOn.After(x.lastDoneOn)
			}
		}
		if x.name != y.name {
			return x.name < y.name
		}
		return false
	}})
}

// printStats prints table with statistics about katas you've done.
func printStats(stats []statistics, showAll, countSort *bool) {
	const format = "%-49v\t%17v\t%10v\n"

	// Print header.
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Kata", "Last done", "Count")
	fmt.Fprintf(tw, format, "----", "---------", "-----")

	// Print lines.
	var katasCount int
	var totalCount int
	sortStats(stats, countSort)
	for _, s := range stats {
		if !show(s, 14, showAll) {
			continue
		}

		katasCount++
		totalCount += s.count

		fmt.Fprintf(tw, format, s.name, formatLastDoneOn(s.lastDoneOn), s.count)
	}
	tw.Flush() // calculate column widths and print table

	// Print footer.
	fmt.Printf("%-49s %30s\n", "----", "-----")
	fmt.Printf("%-49d %30d\n", katasCount, totalCount)
}

// show decides when to show a kata statistics.
func show(s statistics, lastDoneDaysAgo float64, showAll *bool) bool {
	if *showAll {
		return true
	}
	return time.Since(s.lastDoneOn).Hours() < 24*lastDoneDaysAgo
}

func formatLastDoneOn(lastDoneOn time.Time) string {
	daysAgo := int(time.Since(lastDoneOn).Hours() / 24)
	weekday := lastDoneOn.Weekday().String()[:3]
	var s string
	if daysAgo > 14 {
		s = fmt.Sprintf("%s (%s)", lastDoneOn.Format("2006-01-02"), weekday)
	} else {
		w := "day"
		if daysAgo != 1 {
			w += "s"
		}
		s = fmt.Sprintf("%d %s ago (%s)", daysAgo, w, weekday)
	}
	return s
}

// getStats creates statistics about katas you've done.
func getStats(katas []kata) []statistics {
	count := make(map[string]int)
	lastDoneOn := make(map[string]time.Time)
	for _, k := range katas {
		if _, ok := lastDoneOn[k.name]; !ok {
			count[k.name] = 1
			lastDoneOn[k.name] = k.doneOn
			continue
		}
		count[k.name]++
		if k.doneOn.After(lastDoneOn[k.name]) {
			lastDoneOn[k.name] = k.doneOn
		}
	}
	var stats []statistics
	for name := range count {
		stats = append(stats, statistics{name, count[name], lastDoneOn[name]})
	}
	return stats
}

// parseFile extracts katas from a file.
func parseFile(filename string) ([]kata, error) {
	var katas []kata

	f, err := os.Open(filename)
	if err != nil {
		return katas, err
	}

	// Regexes
	kataLineRE := regexp.MustCompile(`^\*\s*([0-9]{4}\-[0-9]{2}\-[0-9]{2}):\s*(.+)$`)
	comaRE := regexp.MustCompile(`\s*,\s*`)

	s := bufio.NewScanner(f)
	for s.Scan() {
		lineParts := kataLineRE.FindStringSubmatch(s.Text())
		if lineParts == nil {
			continue
		}
		date, katasStr := lineParts[1], lineParts[2]

		for _, name := range comaRE.Split(katasStr, -1) {
			if name == "" {
				continue
			}
			doneOn, err := time.Parse("2006-01-02", date)
			if err != nil {
				return katas, err
			}
			katas = append(katas, kata{name, doneOn})
		}
	}
	if s.Err() != nil {
		return katas, s.Err()
	}

	return katas, nil
}
