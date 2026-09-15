package main

func isIsomorphic(s string, t string) bool {
	m := make(map[byte]byte)
	mm := make(map[byte]byte)

	for k, v := range s {
		if _, ok := m[t[k]]; !ok {
			m[t[k]] = byte(v)
		} else {
			if m[t[k]] != byte(v) {
				return false
			}
		}

		if _, ok := mm[byte(v)]; !ok {
			mm[byte(v)] = t[k]
		} else {
			if mm[byte(v)] != t[k] {
				return false
			}
		}

	}

	return true
}
