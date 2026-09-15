package main

func longestPalindrome3(words []string) int {
	mm := make(map[string]int)
	res := 0
	for _, w := range words {
		s := string(w[1]) + string(w[0])
		if mm[s] > 0 {
			mm[s]--
			res += 4
		} else {
			mm[w]++
		}
	}

	for w, v := range mm {
		if w[0] == w[1] && v > 0 {
			res += 2
			break
		}
	}

	return res
}
