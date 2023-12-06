package main

import (
	"flag"
	"os"

	greputil "github.com/heckphy/wb-l2/develop/dev05/grep"
)

func main() {
	opts := greputil.Options{}

	flag.IntVar(&opts.After, "A", 0, "\"after\" печатать +N строк после совпадения")
	flag.IntVar(&opts.Before, "B", 0, "\"before\" печатать +N строк до совпадения")
	flag.IntVar(&opts.Context, "C", 0, "\"context\" (A+B) печатать ±N строк вокруг совпадения")
	flag.BoolVar(&opts.Count, "c", false, "\"count\" (количество строк)")
	flag.BoolVar(&opts.IgnoreCase, "i", false, "\"ignore-case\" (игнорировать регистр)")
	flag.BoolVar(&opts.Invert, "v", false, "\"invert\" (вместо совпадения, исключать)")
	flag.BoolVar(&opts.Fixed, "F", false, "\"fixed\", точное совпадение со строкой, не паттерн")
	flag.BoolVar(&opts.LineNum, "n", false, "\"line num\", напечатать номер строки")
	flag.Parse()

	pattern := flag.Arg(0)
	filename := flag.Arg(1)

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	grep := greputil.NewGrep(pattern, file, os.Stdout, opts)

	grep.Run()
}
