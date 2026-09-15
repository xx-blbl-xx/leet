package main

func isValid20(s string) bool {
	l := make([]byte, 0)
	m := map[byte]byte{
		'[': ']',
		'(': ')',
		'{': '}',
	}
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			l = append(l, s[i])
		} else {
			if len(l) == 0 {
				return false
			}
			if m[l[len(l)-1]] != s[i] {
				return false
			}
			l = l[:len(l)-1]
		}
	}

	return len(l) == 0
}
