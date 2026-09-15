package main

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}

	m := make(map[byte]int)
	mm := make(map[byte]int)
	for k, v := range p {
		m[byte(v)]++
		mm[byte(s[k])]++
	}

	i := 0
	j := len(p) - 1
	res := []int{}
	for j < len(s) {

		r := 0
		for k, v := range mm {
			t := m[k] - v
			if t < 0 {
				t = -t
			}
			r += t
		}
		if r == 0 {
			res = append(res, i)
			r = 1
		}

		mm[s[i]]--
		i++
		j++
		if j == len(s) {
			break
		}
		mm[s[j]]++
	}

	return res
}
