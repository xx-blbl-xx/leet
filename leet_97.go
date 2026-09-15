package main

func isInterleave(s1 string, s2 string, s3 string) bool {
	if s1 == "" {
		return s2 == s3
	}
	if s2 == "" {
		return s1 == s3
	}

	l1 := len(s1)
	l2 := len(s2)

	if l1+l2 != len(s3) {
		return false
	}

	res := make([][]bool, l1+1)

	for k := range res {
		res[k] = make([]bool, l2+1)
	}
	res[0][0] = true

	for i := 1; i <= l1; i++ {
		if s1[i-1] != s3[i-1] {
			break
		}
		res[i][0] = true
	}

	for j := 1; j <= l2; j++ {
		if s2[j-1] != s3[j-1] {
			break
		}
		res[0][j] = true

	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			res[i][j] = (res[i-1][j] && s1[i-1] == s3[i+j-1]) ||
				(res[i][j-1] && s2[j-1] == s3[i+j-1])
		}
	}

	return res[l1][l2]
}
