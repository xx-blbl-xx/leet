package main

func percentageLetter(s string, letter byte) int {
	m := make(map[byte]int, 0)
	for _, v := range s {
		m[byte(v)]++
	}
	return m[letter] * 100 / len(s)
}
