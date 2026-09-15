package main

func permute46(nums []int) [][]int {
	res := [][]int{}

	for i := 0; i < len(nums); i++ {
		res = makePremute(nums, res)
	}

	return res
}

func makePremute(nums []int, r [][]int) [][]int {
	res := [][]int{}
	for _, v := range nums {
		if len(r) == 0 {
			res = append(res, []int{v})
		} else {
			for _, n := range r {
				f := false
				t := []int{}
				for _, d := range n {
					if d == v {
						f = true
						break
					}
					t = append(t, d)
				}
				if f {
					continue
				}
				t = append(t, v)
				res = append(res, t)
			}
		}
	}

	return res
}
