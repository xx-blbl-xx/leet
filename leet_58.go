package main

import "strings"

func lengthOfLastWord(s string) int {
	strs := strings.Split(s, " ")
	l := len(strs) - 1
	for ; l >= 0; l-- {
		if len(strs[l]) != 0 {
			return len(strs[l])
		}
	}

	return 0
}
