package greputil

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
)

type Options struct {
	After      int
	Before     int
	Context    int
	Count      bool
	IgnoreCase bool
	Invert     bool
	Fixed      bool
	LineNum    bool
}

type Grep struct {
	pattern string
	r       io.Reader
	w       io.Writer
	opts    Options
}

func NewGrep(pattern string, r io.Reader, w io.Writer, opts Options) *Grep {
	if opts.Context != 0 {
		opts.Before = opts.Context
		opts.After = opts.Context
	}

	if opts.Fixed {
		pattern = regexp.QuoteMeta(pattern)
	}

	if opts.IgnoreCase {
		pattern = "(?i)" + pattern
	}

	return &Grep{
		pattern,
		r,
		w,
		opts,
	}
}

func (g *Grep) Run() error {
	if !g.opts.Count {
		return g.parse()
	}

	cnt, err := g.count()
	if err != nil {
		return err
	}
	fmt.Fprintln(g.w, cnt)

	return nil
}

func (g *Grep) parse() error {
	r := bufio.NewReader(g.r)
	beforeBuf := make([]string, g.opts.Before)
	var after int
	var i int
	var lastI int

	for {
		line, err := r.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				break
			}

			return err
		}

		i++

		if after > 0 {
			if !g.matches(line) {
				g.print(line, i, true)
				lastI = i
				after--
				continue
			}

			after = 0
		}

		if g.matches(line) {
			if (g.opts.Before != 0 || g.opts.After != 0) && i-lastI > 1 && lastI != 0 {
				fmt.Fprintln(g.w, "--")
			}

			offset := i - g.opts.Before

			for j, prevLine := range beforeBuf {
				g.print(prevLine, j+offset, true)
			}

			g.print(line, i, false)
			lastI = i

			beforeBuf = beforeBuf[:0]
			after = g.opts.After
			continue
		}

		if g.opts.Before > 0 {
			if len(beforeBuf) == g.opts.Before {
				beforeBuf = beforeBuf[1:]
			}

			beforeBuf = append(beforeBuf, line)
		}
	}

	return nil
}

func (g *Grep) count() (int, error) {
	r := bufio.NewReader(g.r)
	var cnt int

	for {
		line, err := r.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				break
			}

			return 0, err
		}

		if g.matches(line) {
			cnt++
		}
	}

	return cnt, nil
}

func (g *Grep) print(line string, lineno int, isCtx bool) {
	if g.opts.Count {
		return
	}

	sep := ':'

	if isCtx {
		sep = '-'
	}

	if g.opts.LineNum {
		fmt.Fprintf(g.w, "%d%c%s", lineno, sep, line)
	} else {
		fmt.Fprint(g.w, line)
	}
}

func (g *Grep) matches(line string) bool {
	re := regexp.MustCompile(g.pattern)
	result := re.MatchString(line)

	if g.opts.Invert {
		result = !result
	}

	return result
}
