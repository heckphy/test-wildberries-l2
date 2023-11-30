package main

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"fmt"
	"sort"
	"strings"
)

func sortString(s string) string {
	runes := []rune(s)

	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	return string(runes)
}

func findAnagrams(words []string) map[string][]string {
	anagrams := make(map[string][]string)
	sets := make(map[string][]string)
	seen := make(map[string]struct{})

	for _, word := range words {
		word = strings.ToLower(word)

		if _, ok := seen[word]; ok {
			continue
		}

		sortedWord := sortString(word)
		sets[sortedWord] = append(sets[sortedWord], word)
		seen[sortedWord] = struct{}{}
	}

	for _, value := range sets {
		if len(value) == 1 {
			continue
		}

		anagrams[value[0]] = value

		sort.Slice(value, func(i, j int) bool {
			return value[i] < value[j]
		})
	}

	return anagrams
}

func main() {
	words := []string{"тяпка", "пятак", "пятка", "листок", "слиток", "столик"}
	anagrams := findAnagrams(words)

	for key, value := range anagrams {
		fmt.Printf("%s: %v\n", key, value)
	}
}
