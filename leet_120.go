package main

func minimumTotal(triangle [][]int) int {

	li := len(triangle)
	if li == 1 {
		return triangle[0][0]
	}

	res := make([][]int, li)
	res[0] = []int{triangle[0][0]}
	res[1] = []int{triangle[0][0] + triangle[1][0], triangle[0][0] + triangle[1][1]}

	for i := 2; i < li; i++ {

		res[i] = make([]int, i+1)
		res[i][0] = res[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			res[i][j] = min(res[i-1][j]+triangle[i][j], res[i-1][j-1]+triangle[i][j])
		}
		res[i][i] = res[i-1][i-1] + triangle[i][i]
	}

	rr := res[li-1][0]
	for _, v := range res[li-1] {
		rr = min(v, rr)
	}
	return rr

}
