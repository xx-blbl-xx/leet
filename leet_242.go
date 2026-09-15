package main

func isAnagram(s string, t string) bool {
	if len(t) != len(s) {
		return false
	}

	ms := map[byte]int{}
	mt := map[byte]int{}

	for _, v := range s {
		ms[byte(v)]++
	}
	for _, v := range t {
		mt[byte(v)]++
	}

	for k, v := range ms {
		if mt[k] != v {
			return false
		}
	}

	return true
}
