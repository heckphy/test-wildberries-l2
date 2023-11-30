package main

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

		if _, ok := sets[sortedWord]; ok {
			sets[sortedWord] = append(sets[sortedWord], word)
		} else {
			sets[sortedWord] = []string{word}
		}

		seen[sortedWord] = struct{}{}
	}

	for _, value := range sets {
		if len(value) == 1 {
			continue
		}

		sort.Slice(value, func(i, j int) bool {
			return value[i] < value[j]
		})

		anagrams[value[0]] = value
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
