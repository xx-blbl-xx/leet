package main

func rowAndMaximumOnes(mat [][]int) []int {
	res := make([][]int, 0, len(mat))

	for k, m := range mat {
		r := 0
		for _, v := range m {
			r += v
		}
		res = append(res, []int{k, r})
	}

	m := -1
	_k := 0
	for k, v := range res {
		if m < v[1] {
			m = v[1]
			_k = k
		}
	}

	return res[_k]
}
