package main

func longestPalindrome2(s string) string {
	ll := len(s)
	if ll == 1 {
		return s
	}

	start, end := 0, 0
	for i := 0; i < ll; i++ {
		s1, e1 := lp2(s, ll, i, i)
		if end-start < e1-s1 {
			start, end = s1, e1
		}

		if i+1 < ll && s[i] == s[i+1] {
			s2, e2 := lp2(s, ll, i, i+1)
			if end-start < e2-s2 {
				start, end = s2, e2
			}
		}

	}

	return s[start : end+1]
}

func lp2(str string, ll, s, e int) (int, int) {
	for str[s] == str[e] {
		s--
		e++
		if s < 0 {
			break
		}
		if e >= ll {
			break
		}
	}

	s++
	e--

	return s, e
}
