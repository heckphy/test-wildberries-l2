package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	unpacked, err := UnpackStr(`qwe\\\`)
	fmt.Println(unpacked, err)
}

func UnpackStr(s string) (string, error) {
	var c rune
	var n strings.Builder
	var result strings.Builder
	var isEscaping bool

	for _, r := range s {
		if unicode.IsDigit(r) && c == 0 {
			return "", errors.New("invalid input string")
		}

		if r == '\\' && !isEscaping {
			isEscaping = true
			continue
		}

		if unicode.IsDigit(r) && !isEscaping {
			n.WriteRune(r)
			continue
		}

		if c == 0 {
			c = r
			continue
		}

		multiplyRune(&result, &n, c)

		isEscaping = false
		c = r
		n.Reset()
	}

	if c != 0 {
		multiplyRune(&result, &n, c)
	}

	if isEscaping {
		return "", errors.New("invalid input string")
	}

	return result.String(), nil
}

func multiplyRune(dst *strings.Builder, n *strings.Builder, r rune) {
	cnt, err := strconv.Atoi(n.String())
	if err != nil {
		cnt = 1
	}

	for i := 0; i < cnt; i++ {
		dst.WriteRune(r)
	}
}
