package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Options struct {
	fields    string
	delimiter string
	separated bool
}

func main() {
	opts := Options{}

	flag.StringVar(&opts.fields, "f", "", "выбрать поля (колонки)")
	flag.StringVar(&opts.delimiter, "d", "\t", "использовать другой разделитель")
	flag.BoolVar(&opts.separated, "s", false, "только строки с разделителем")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	lines := []string{}

	for {
		fmt.Print(">> ")
		line, _ := reader.ReadString('\n')

		if line == "\n" || line == "\r\n" {
			break
		}

		line = strings.TrimRight(line, "\r\n")

		lines = append(lines, line)
		fmt.Println(line)
	}

	res := []string{}

	for _, line := range lines {
		fields := strings.Split(line, opts.delimiter)

		if opts.separated && len(fields) == 1 {
			continue
		}

		fields, _ = filterFields(fields, opts.fields)

		res = append(res, strings.Join(fields, opts.delimiter))
	}

	for _, r := range res {
		fmt.Println(r)
	}
}

func filterFields(fields []string, crit string) (res []string, err error) {
	if len(fields) == 1 {
		return fields, nil
	}

	poses := strings.Split(crit, ",")

	for _, pos := range poses {
		rng := strings.SplitN(pos, "-", 2)

		if len(rng) == 1 {
			i, _ := strconv.Atoi(rng[0])
			res = append(res, fields[i-1])
			continue
		}

		from, err := strconv.Atoi(rng[0])
		if err != nil {
			from = 1
		}

		to, err := strconv.Atoi(rng[1])
		if err != nil {
			to = len(fields)
		}

		for i := from; i <= to; i++ {
			res = append(res, fields[i-1])
		}
	}

	return res, err
}
