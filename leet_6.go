package main

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	mm := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		mm[i] = make([]byte, len(s))
	}
	t := 0
	l := len(s)
	for i := 0; i < l; {
		for j := 0; j < numRows && i < l; j++ {
			mm[j][t*(numRows-1)] = s[i]
			i++
		}
		for j := numRows - 2; j > 0 && i < l; j-- {
			mm[j][t*(numRows-1)+j] = s[i]
			i++
		}
		t++
	}

	res := make([]byte, 0, len(s))
	for i := 0; i < numRows; i++ {
		for j := 0; j < l; j++ {
			if mm[i][j] == 0 {
				continue
			}
			res = append(res, mm[i][j])
		}
	}

	return string(res)
}
