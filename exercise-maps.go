package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, word := range strings.Fields(s) {
		_, word_found := m[word]
		if word_found == true {
			m[strings.ToLower(word)] += 1
		} else {
			m[strings.ToLower(word)] = 1
		}
	}
	return m
}

func main() {
	fmt.Println(WordCount("The quick brown fox jumped over the lazy dog."))
}
