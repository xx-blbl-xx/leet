package main

func canConstruct(ransomNote string, magazine string) bool {
	m := map[byte]int{}

	for _, v := range magazine {
		m[byte(v)]++
	}

	for _, v := range ransomNote {
		m[byte(v)]--
	}

	for _, v := range m {
		if v < 0 {
			return false
		}
	}

	return true
}
