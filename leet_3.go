package main

func lengthOfLongestSubstring3(s string) int {
	if len(s) == 0 {
		return 0
	}
	m := make(map[byte]int)
	res := 1
	m[s[0]] = 1
	i, j := 0, 1
	for i < j && j < len(s) {

		for j < len(s) {
			if _, ok := m[s[j]]; ok {
				res = max(res, j-i)
				break
			}

			m[s[j]] = 1
			j++
		}

		if j == len(s) {
			break
		}

		for ; i < j; i++ {
			if s[i] == s[j] {
				i++
				j++
				break
			}
			delete(m, s[i])
		}

	}

	return max(res, len(m))
}
