package main

func findBall(grid [][]int) []int {
	m := len(grid)
	n := len(grid[0])
	res := make([]int, n)

	for i := 0; i < n; i++ {

		k := i
		for j := 0; j < m; j++ {
			if grid[j][k] == 1 {
				if k+1 >= n || grid[j][k+1] == -1 {
					k = -1
					break
				}
				k += 1
			} else {
				if k-1 < 0 || grid[j][k-1] == 1 {
					k = -1
					break
				}
				k -= 1
			}

		}

		res[i] = k
	}

	return res
}
