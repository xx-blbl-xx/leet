package main

import "strings"

func reverseWords(s string) string {
	strsT := strings.Split(s, " ")
	strs := make([]string, 0, len(strsT))
	for i := len(strsT) - 1; i >= 0; i-- {
		if strsT[i] == "" {
			continue
		}
		strs = append(strs, strsT[i])
	}

	return strings.Join(strs, " ")
}
