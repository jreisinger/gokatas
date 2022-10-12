package gokatas

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/tabwriter"
	"time"
)

// KatasFile is a MarkDown file to track katas you've done. It looks like this:
//
//   - 2022-09-13: boring/boring, boring/channel
//   - 2022-09-10: areader
const KatasFile = "katas.md"

// Kata represents a programming kata.
type Kata struct {
	Name      string
	LastDone  time.Time
	TimesDone int
	Level     string
	Topics    []string
}

// Get returns all existing katas and your practice statistics.
func Get() ([]Kata, error) {
	existing, err := getExisting()
	if err != nil {
		return nil, err
	}

	done, err := getDone()
	if err != nil {
		return nil, err
	}

HERE:
	for _, d := range done {
		for _, e := range existing {
			if d.Name == e.Name {
				continue HERE
			}
		}
		log.Printf("kata '%s' stated in %s does not exist in this repo", d.Name, KatasFile)
	}

	for i := range existing {
		for j := range done {
			if existing[i].Name == done[j].Name {
				existing[i].TimesDone = done[j].TimesDone
				existing[i].LastDone = done[j].LastDone
			}
		}
	}

	return existing, nil
}

// getExistings returns all existing katas.
func getExisting() ([]Kata, error) {
	cmd := exec.Command("go", "list", "-f", "{{.Dir}}", "./...")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	var katas []Kata
	for _, line := range strings.Split(string(out), "\n") {
		name := strings.TrimPrefix(line, cwd)
		name = strings.TrimPrefix(name, "/")
		if name == "" || strings.HasSuffix(name, "cmd") {
			continue
		}
		level, topics, err := parseKata(name)
		if err != nil {
			return nil, err
		}
		katas = append(katas, Kata{Name: name, Level: level, Topics: uniq(topics)})
	}
	return katas, err
}

// uniq removes duplicates from topics.
func uniq(topics []string) []string {
	seen := make(map[string]bool)
	var unique []string
	for _, topic := range topics {
		if _, ok := seen[topic]; !ok {
			seen[topic] = true
			unique = append(unique, topic)
		}
	}
	return unique
}

// getDone returns katas from the KatasFile.
func getDone() ([]Kata, error) {
	f, err := os.Open(KatasFile)
	if err != nil {
		return nil, err
	}

	// Regexes
	kataLineRE := regexp.MustCompile(`^\s*\*\s*([0-9]{4}\-[0-9]{2}\-[0-9]{2}):\s*(.+)$`)
	comaRE := regexp.MustCompile(`\s*,\s*`) // works both with w1,w2 and w1, w2

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

			name = strings.TrimSpace(name)

			if kata, ok := katas[name]; ok {
				kata.TimesDone++
				if doneOn.After(kata.LastDone) {
					kata.LastDone = doneOn
				}
				katas[name] = kata
			} else {
				kata.Name = name
				kata.TimesDone = 1
				kata.LastDone = doneOn
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

func parseKata(name string) (level string, topics []string, err error) {
	fn := func(path string, d fs.DirEntry, err error) error {
		if filepath.Ext(path) == ".go" {
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			defer f.Close()
			s := bufio.NewScanner(f)
			for s.Scan() {
				line := s.Text()
				if strings.HasPrefix(line, "// Level:") {
					level = grepLevel(s.Text())
				}
				if strings.HasPrefix(line, "// Topics:") {
					topics = append(topics, grepTopics(s.Text())...)
				}
			}
			if err := s.Err(); err != nil {
				return err
			}
		}
		return nil
	}
	absPath, err := filepath.Abs(name)
	if err != nil {
		return "", nil, err
	}
	err = filepath.WalkDir(absPath, fn)
	if err != nil {
		return "", nil, err
	}
	return level, topics, err
}

func grepLevel(line string) string {
	_, level, _ := strings.Cut(line, ":")
	return strings.TrimSpace(level)
}

func grepTopics(line string) []string {
	_, topicsStr, _ := strings.Cut(line, ":")
	topics := strings.Split(topicsStr, ",")
	for i := range topics {
		topics[i] = strings.TrimSpace(topics[i])
	}
	return topics
}

// Print prints table with statistics about katas. Only katas of level (if not
// empty) and lastDoneDaysAgo or sooner are shown. Katas are sorted by column.
func Print(katas []Kata, lastDoneDaysAgo int, column int, level string) {
	const format = "%v\t%v\t%5v\t%v\t%v\n"

	// Print header.
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Kata", "Last done", "Done", "Level", "Topics")
	fmt.Fprintf(tw, format, "----", "---------", "----", "-----", "------")

	// Print lines.
	var katasCount int
	var totalCount int
	sortKatas(katas, &column)
	for _, k := range katas {
		if !show(k, lastDoneDaysAgo) {
			continue
		}
		if level != "" && k.Level != level {
			continue
		}

		katasCount++
		totalCount += k.TimesDone

		fmt.Fprintf(tw, format, k.Name, humanize(k.LastDone), fmt.Sprintf("%dx", k.TimesDone), k.Level, strings.Join(k.Topics, ", "))
	}
	// Print footer.
	fmt.Fprintf(tw, format, "----", "", "----", "", "")
	fmt.Fprintf(tw, format, katasCount, "", totalCount, "", "")

	tw.Flush() // calculate column widths and print table
}

type customSort struct {
	katas []Kata
	less  func(x, y Kata) bool
}

func (x customSort) Len() int           { return len(x.katas) }
func (x customSort) Less(i, j int) bool { return x.less(x.katas[i], x.katas[j]) }
func (x customSort) Swap(i, j int)      { x.katas[i], x.katas[j] = x.katas[j], x.katas[i] }

// sortKatas sorts katas by column. Not all columns are sortable. Secondary sort
// orders is always by kata name.
func sortKatas(katas []Kata, column *int) {
	sort.Sort(customSort{katas, func(x, y Kata) bool {
		switch *column {
		case 1:
			if x.Name != y.Name {
				return x.Name < y.Name
			}
		case 2:
			if x.LastDone != y.LastDone {
				return x.LastDone.After(y.LastDone)
			}
		case 3:
			if x.TimesDone != y.TimesDone {
				return x.TimesDone > y.TimesDone
			}
		default:
			log.Fatalf("can't sort by column %d", *column)
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
	return k.LastDone.After(t)
}

// humanize make the time easier to read for humans.
func humanize(lastDone time.Time) string {
	if lastDone.IsZero() {
		return "never"
	}
	daysAgo := int(time.Since(lastDone).Hours() / 24)
	w := "day"
	if daysAgo != 1 {
		w += "s"
	}
	return fmt.Sprintf("%d %s ago", daysAgo, w)
}
