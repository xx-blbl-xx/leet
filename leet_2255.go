package main

import "strings"

func countPrefixes(words []string, s string) int {
	res := 0
	for _, w := range words {
		if strings.HasPrefix(s, w) {
			res++
		}
	}
	return res
}
