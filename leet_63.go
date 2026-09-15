package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	if obstacleGrid[m-1][n-1] == 1 || obstacleGrid[0][0] == 1 {
		return 0
	}

	if m == 1 {
		for _, v := range obstacleGrid[0] {
			if v == 1 {
				return 0
			}
		}
		return 1
	}
	if n == 1 {
		for _, v := range obstacleGrid {
			if v[0] == 1 {
				return 0
			}
		}
		return 1
	}

	res := make([][]int, m)
	for k := range res {
		res[k] = make([]int, n)
	}
	for k := range res {
		if obstacleGrid[k][0] == 1 {
			break
		} else {
			res[k][0] = 1
		}
	}
	for k := range res[0] {
		if obstacleGrid[0][k] == 1 {
			break
		} else {
			res[0][k] = 1
		}
	}

	for i := 1; i < m; i++ {

		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				res[i][j] = 0
			} else {
				res[i][j] = res[i-1][j] + res[i][j-1]
			}
		}
	}

	return res[m-1][n-1]
}
