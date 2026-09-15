package main

import "math"

func maximalSquare(matrix [][]byte) int {
	l1 := len(matrix)
	l2 := len(matrix[0])

	res := make([][]int, l1)
	for k := range res {
		res[k] = make([]int, l2)
	}

	maxres := 0
	for i := 0; i < l1; i++ {
		if matrix[i][0] == '1' {
			res[i][0] = 1
		}
		maxres = max(maxres, res[i][0])
	}
	for j := 0; j < l2; j++ {
		if matrix[0][j] == '1' {
			res[0][j] = 1
		}
		maxres = max(maxres, res[0][j])
	}

	for i := 1; i < l1; i++ {

		for j := 1; j < l2; j++ {
			if matrix[i][j] == '0' {
				res[i][j] = 0
			} else {
				if res[i-1][j-1] > 0 && res[i-1][j] > 0 && res[i][j-1] > 0 {
					t := int(min(math.Sqrt(float64(res[i-1][j-1])), math.Sqrt(float64(res[i][j-1])), math.Sqrt(float64(res[i-1][j]))))
					res[i][j] = (t + 1) * (t + 1)
				} else {
					res[i][j] = 1
				}
			}

			maxres = max(maxres, res[i][j])
		}
	}

	return maxres
}
