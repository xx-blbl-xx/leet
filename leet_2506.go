package main

func similarPairs(words []string) int {
	mmp := make([]map[byte]bool, 0, len(words))
	for _, word := range words {
		m := make(map[byte]bool)
		for _, w := range word {
			m[byte(w)] = true
		}
		mmp = append(mmp, m)
	}

	res := 0
	ll := len(mmp)
	for i := 0; i < ll; i++ {
		for j := i + 1; j < ll; j++ {
			if checkmmp(mmp[i], mmp[j]) {
				res++
			}
		}
	}

	return res
}

func checkmmp(mm1, mm2 map[byte]bool) bool {
	if len(mm1) != len(mm2) {
		return false
	}

	for k := range mm1 {
		if !mm2[k] {
			return false
		}
	}

	return true
}
