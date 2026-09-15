package main

func uniquePaths62(m int, n int) int {

	res := make([][]int, m)
	for k := range res {
		res[k] = make([]int, n)
	}
	for k := range res {
		res[k][0] = 1
	}
	for k := range res[0] {
		res[0][k] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			res[i][j] = res[i-1][j] + res[i][j-1]
		}
	}

	return res[m-1][n-1]
}
