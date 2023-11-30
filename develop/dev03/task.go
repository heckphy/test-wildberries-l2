package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"time"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Options struct {
	Key                 int
	NumericSort         bool
	Reverse             bool
	Unique              bool
	MonthSort           bool
	IgnoreLeadingBlanks bool
	Check               bool
	HumanNumericSort    bool
}

var months = map[string]time.Month{
	"jan": time.January,
	"feb": time.February,
	"mar": time.March,
	"apr": time.April,
	"may": time.May,
	"jun": time.June,
	"jul": time.July,
	"aug": time.August,
	"sep": time.September,
	"oct": time.October,
	"nov": time.November,
	"dec": time.December,
}

func unique(slice []string) []string {
	seen := make(map[string]struct{})
	set := make([]string, len(slice))

	for _, entry := range slice {
		if _, exists := seen[entry]; !exists {
			seen[entry] = struct{}{}
			set = append(set, entry)
		}
	}

	return set
}

func readLines(filename string) (lines []string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := bufio.NewReader(file)

	for {
		const delim = '\n'

		line, err := r.ReadString(delim)

		if err == nil || len(line) > 0 {
			if err != nil {
				line += string(delim)
			}

			lines = append(lines, line)
		}

		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, err
		}
	}

	return lines, nil
}

func writeLines(filename string, lines []string) (err error) {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	defer w.Flush()

	for _, line := range lines {
		_, err := w.WriteString(line)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	opts := Options{}

	flag.IntVar(&opts.Key, "k", 0, "указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)")
	flag.BoolVar(&opts.NumericSort, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&opts.Reverse, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&opts.Unique, "u", false, "не выводить повторяющиеся строки")
	flag.BoolVar(&opts.MonthSort, "M", false, "сортировать по названию месяца")
	flag.BoolVar(&opts.IgnoreLeadingBlanks, "b", false, "игнорировать хвостовые пробелы")
	flag.BoolVar(&opts.Check, "c", false, "проверять отсортированы ли данные")
	flag.BoolVar(&opts.HumanNumericSort, "h", false, "сортировать по числовому значению с учетом суффиксов")
	flag.Parse()

	lines, err := readLines("lines.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	less := func(i, j int) bool {
		var a string
		var b string

		if opts.Key != 0 {
			fields1 := strings.Split(lines[i], " ")
			fields2 := strings.Split(lines[j], " ")

			field1 := fields1[opts.Key-1]
			field2 := fields2[opts.Key-1]

			a = field1
			b = field2
		} else {
			a = lines[i]
			b = lines[j]
		}

		if opts.NumericSort {
			var n1 int
			fmt.Sscanf(a, "%d", &n1)

			var n2 int
			fmt.Sscanf(b, "%d", &n2)

			return n1 < n2
		} else if opts.MonthSort {
			month1 := strings.ToLower(a[:3])
			month2 := strings.ToLower(b[:3])

			if opts.IgnoreLeadingBlanks {
				month1 = strings.TrimLeft(month1, " ")
				month2 = strings.TrimLeft(month2, " ")
			}

			m1 := months[month1]
			m2 := months[month2]

			return m1 < m2
		} else if opts.HumanNumericSort {
			var n1, n2 int
			var s1, s2 string
			fmt.Sscanf(a, "%d%s", &n1, &s1)
			fmt.Sscanf(b, "%d%s", &n2, &s2)

			if s1 == s2 {
				return n1 < n2
			}

			return s1 < s2
		} else {
			if opts.IgnoreLeadingBlanks {
				a = strings.TrimLeft(a, " ")
				b = strings.TrimLeft(b, " ")
			}

			return a < b
		}
	}

	if opts.Check {
		for i, length := 0, len(lines); i < length-1; i++ {
			if !less(i, i+1) {
				fmt.Printf("sort: lines.txt:%d: disorder: %s", i+2, lines[i+1])
				break
			}
		}

		return
	}

	sort.Slice(lines, less)

	if opts.Unique {
		lines = unique(lines)
	}

	err = writeLines("sorted_lines.txt", lines)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
