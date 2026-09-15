package main

func letterCombinations17(digits string) []string {
	m := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	res := m[digits[0]]

	for k, v := range digits {
		if k == 0 {
			continue
		}
		res = lettr17(m, byte(v), res)
	}

	return res
}

func lettr17(m map[byte][]string, d byte, rs []string) []string {
	res := []string{}

	for _, v := range m[d] {
		for _, w := range rs {
			res = append(res, w+v)
		}
	}

	return res
}
