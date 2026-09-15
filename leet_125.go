package main

import "strings"

func isPalindrome125(s string) bool {
	t := strings.ToLower(s)
	ss := make([]rune, 0)
	for _, v := range t {
		if (v >= '0' && v <= '9') || (v >= 'a' && v <= 'z') {
			ss = append(ss, v)
		}
	}
	s = string(ss)

	i := 0
	j := len(s) - 1

	var ii, jj byte
	for i <= j {
		ii = s[i]
		jj = s[j]

		if ii != jj {
			return false
		}
		i++
		j--
	}

	return true
}
