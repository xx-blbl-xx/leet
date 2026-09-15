package main

func minPathSum64(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	if m == 1 {
		r := 0
		for _, v := range grid[0] {
			r += v
		}
		return r
	}

	if n == 1 {
		r := 0
		for i := 0; i < m; i++ {
			r += grid[i][0]
		}
		return r
	}

	res := make([][]int, m)
	for k := range res {
		res[k] = make([]int, n)
	}

	r := 0
	for i := 0; i < m; i++ {
		r += grid[i][0]
		res[i][0] = r
	}
	r = 0
	for i := 0; i < n; i++ {
		r += grid[0][i]
		res[0][i] = r
	}

	for i := 1; i < m; i++ {

		for j := 1; j < n; j++ {
			res[i][j] = min(res[i-1][j], res[i][j-1]) + grid[i][j]
		}
	}

	return res[m-1][n-1]

}
