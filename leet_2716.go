package main

func minimizedStringLength(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	sm := make(map[byte]bool)
	for _, v := range s {
		sm[byte(v)] = true
	}

	return len(sm)
}
