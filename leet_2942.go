package main

import "strings"

func findWordsContaining(words []string, x byte) []int {
	res := make([]int, 0)
	for i, word := range words {
		if strings.Contains(word, string(x)) {
			res = append(res, i)
		}
	}

	return res
}
