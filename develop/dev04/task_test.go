package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestFindAnagramsExcludeSetsWithLenOfOne(t *testing.T) {
	words := []string{"тяпка", "корыто"}

	anagrams := findAnagrams(words)

	if len(anagrams) > 0 {
		t.Logf(`findAnagram results is %v, expected map[]`, anagrams)
		t.Fail()
	}
}

func TestFindAnagramsAnagramsLowercased(t *testing.T) {
	words := []string{"ТЯПКА", "ПЯТАК"}
	anagrams := findAnagrams(words)

	for key, value := range anagrams {
		if key != strings.ToLower(key) {
			t.Logf(`key in map is %s, expected %s`, key, strings.ToLower(key))
			t.Fail()
		}

		for _, word := range value {
			if word != strings.ToLower(word) {
				t.Logf(`word in set is %s, expected %s`, word, strings.ToLower(word))
				t.Fail()
			}
		}
	}
}

func TestFindAnagramsFirstSetItemIsDictKey(t *testing.T) {
	words := []string{"тяпка", "пЯтак", "пятка"}
	anagrams := findAnagrams(words)
	expected := "тяпка"

	for key := range anagrams {
		if key != expected {
			t.Logf(`key is %s, expected %s`, key, expected)
			t.Fail()
		}
	}
}

func TestAnagramAnalisysSorting(t *testing.T) {
	words := []string{"тяпка", "пЯтак", "пятка"}
	anagrams := findAnagrams(words)
	expected := []string{"пятак", "пятка", "тяпка"}

	for _, value := range anagrams {
		if !reflect.DeepEqual(value, expected) {
			t.Logf(`set is not sorted: %v, expected %v`, value, expected)
			t.Fail()
		}
	}
}
