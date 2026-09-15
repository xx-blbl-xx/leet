package main

import "strings"

func wordPattern(pattern string, s string) bool {
	ss := strings.Split(s, " ")

	if len(pattern) != len(ss) {
		return false
	}

	m := make(map[byte]string)
	mm := make(map[string]byte)
	for k, v := range pattern {
		if _, ok := m[byte(v)]; !ok {
			m[byte(v)] = ss[k]
		} else {
			if m[byte(v)] != ss[k] {
				return false
			}
		}

		if _, ok := mm[ss[k]]; !ok {
			mm[ss[k]] = byte(v)
		} else {
			if mm[ss[k]] != byte(v) {
				return false
			}
		}
	}

	return true
}
