package main

func wordBreak139(s string, wordDict []string) bool {
	res := make([]bool, len(s))

	mw := make(map[string]bool, 0)
	for _, v := range wordDict {
		mw[v] = true
	}

	if mw[string(s[0])] {
		res[0] = true
	} else {
		res[0] = false
	}

	for i := 1; i < len(s); i++ {

		r := false
		if mw[string(s[:i+1])] {
			r = true
		} else {
			for k := 0; k < i; k++ {
				if res[k] && mw[string(s[k+1:i+1])] {
					r = true
					break
				}
			}
		}

		res[i] = r
	}

	return res[len(s)-1]
}
