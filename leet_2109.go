package main

func addSpaces(s string, spaces []int) string {
	res := make([]byte, 0, len(s)+len(spaces))
	j := 0
	for i := 0; i < len(s); i++ {
		if j < len(spaces) && spaces[j] == i {
			res = append(res, ' ')
			j++
		}
		res = append(res, s[i])
	}

	return string(res)
}
