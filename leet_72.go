package main

func minDistance72(word1 string, word2 string) int {
	l1 := len(word1)
	l2 := len(word2)

	if l1 == 0 {
		return l2
	}
	if l2 == 0 {
		return l1
	}

	res := make([][]int, l1+1)

	for k := range res {
		res[k] = make([]int, l2+1)
	}

	for i := 0; i <= l1; i++ {
		res[i][0] = i
	}
	for j := 0; j <= l2; j++ {
		res[0][j] = j
	}

	for i := 1; i <= l1; i++ {

		for j := 1; j <= l2; j++ {
			if word1[i-1] == word2[j-1] {
				res[i][j] = 1 + min(res[i-1][j], res[i][j-1], res[i-1][j-1]-1)
			} else {
				res[i][j] = 1 + min(res[i-1][j], res[i][j-1], res[i-1][j-1])
			}

		}
	}

	return res[l1][l2]

}
