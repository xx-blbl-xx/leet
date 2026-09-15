package main

func minWindow(s string, t string) string {
	sm := make(map[byte]int)
	st := make(map[byte]int)

	for i := 0; i < len(t); i++ {
		st[t[i]]++
	}

	l, r := 0, 0
	small := len(s)
	resl, resr := -1, -1
	for i := 0; i < len(s); i++ {
		if st[s[i]] > 0 {
			sm[s[i]]++
			r = i
			l = checkWindow(sm, st, s, t, l, r)

			if checksame(sm, st) {
				if small >= r+1-l {
					small = r + 1 - l
					resl, resr = l, r
				}
			}

		}
	}
	if resl < 0 {
		return ""
	}
	return s[resl : resr+1]
}

func checksame(sm, st map[byte]int) bool {
	for k, v := range st {
		if sm[k] < v {
			return false
		}
	}

	return true
}

func checkWindow(sm, st map[byte]int, s, t string, l, r int) int {
	if r-l < len(t)-1 {
		return l
	}

	for ; l < r-len(t)+1; l++ {
		if st[s[l]] > 0 {
			if sm[s[l]] > st[s[l]] {
				sm[s[l]]--
			} else {
				return l
			}
		}

	}

	return l
}
