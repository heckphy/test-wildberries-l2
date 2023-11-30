package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Options struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
}

type Arguments struct {
	pattern  string
	filename string
}

func matches(pattern string, line string, invert bool) (result bool) {
	re := regexp.MustCompile(pattern)
	result = re.MatchString(line)

	if invert {
		result = !result
	}

	return result
}

func formatLine(line string, i int, lineNum bool, isCtx bool) string {
	var prefix string

	if lineNum {
		prefix = fmt.Sprintf("%d:", i+1)

		if isCtx {
			prefix = fmt.Sprintf("%d-", i+1)
		}
	}

	return fmt.Sprintf("%s%s", prefix, line)
}

func grep(pattern string, lines []string, opts Options) (res []string, cnt int) {
	if opts.fixed {
		pattern = regexp.QuoteMeta(pattern)
	}

	if opts.ignoreCase {
		pattern = "(?i)" + pattern
	}

	seen := make(map[int]struct{}, len(lines))

	for i, line := range lines {
		if !matches(pattern, line, opts.invert) {
			continue
		}

		stepB := opts.before
		stepA := opts.after

		if opts.context != 0 {
			stepB = opts.context
			stepA = opts.context
		}

		if stepB != 0 {
			before := make([]string, 0, stepB)

			for j := max(0, i-stepB); j < i; j++ {
				if matches(pattern, lines[j], opts.invert) {
					before = nil
					continue
				}

				if _, ok := seen[j]; ok {
					continue
				}

				before = append(before, formatLine(lines[j], j, opts.lineNum, true))
				seen[j] = struct{}{}
			}

			res = append(res, before...)
		}

		res = append(res, formatLine(line, i, opts.lineNum, false))
		seen[i] = struct{}{}
		cnt += 1

		if stepA != 0 {
			for j, to := i+1, min(len(lines), i+1+stepA); j < to; j++ {
				if matches(pattern, lines[j], opts.invert) {
					break
				}

				res = append(res, formatLine(lines[j], j, opts.lineNum, true))
				seen[j] = struct{}{}
			}
		}
	}

	return res, cnt
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

func main() {
	opts := Options{}

	flag.IntVar(&opts.after, "A", 0, "\"after\" печатать +N строк после совпадения")
	flag.IntVar(&opts.before, "B", 0, "\"before\" печатать +N строк до совпадения")
	flag.IntVar(&opts.context, "C", 0, "\"context\" (A+B) печатать ±N строк вокруг совпадения")
	flag.BoolVar(&opts.count, "c", false, "\"count\" (количество строк)")
	flag.BoolVar(&opts.ignoreCase, "i", false, "\"ignore-case\" (игнорировать регистр)")
	flag.BoolVar(&opts.invert, "v", false, "\"invert\" (вместо совпадения, исключать)")
	flag.BoolVar(&opts.fixed, "F", false, "\"fixed\", точное совпадение со строкой, не паттерн")
	flag.BoolVar(&opts.lineNum, "n", false, "\"line num\", напечатать номер строки")
	flag.Parse()

	args := Arguments{}

	args.pattern = flag.Arg(0)
	args.filename = flag.Arg(1)

	lines, err := readLines(args.filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	result, cnt := grep(args.pattern, lines, opts)

	fmt.Println(result)
	fmt.Println(cnt)
}
