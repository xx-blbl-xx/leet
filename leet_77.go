package main

func combine(n int, k int) [][]int {
	res := [][]int{}

	for i := 0; i < k; i++ {
		res = makeCombine(i+1, n, res)
	}

	return res
}

func makeCombine(i, j int, r [][]int) [][]int {
	res := [][]int{}
	for i <= j {
		if len(r) == 0 {
			res = append(res, []int{i})
		} else {
			for _, v := range r {
				x := v[len(v)-1]
				if x >= i {
					continue
				}
				vt := append([]int{}, v...)
				vt = append(vt, i)
				res = append(res, vt)
			}
		}

		i++
	}
	return res
}
