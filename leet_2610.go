package main

func findMatrix(nums []int) [][]int {
	res := make([][]int, 0)
	res = append(res, []int{})
	rt := make([]map[int]bool, 0)
	rt = append(rt, map[int]bool{})

	for i := 0; i < len(nums); i++ {
		flag := false
		for k, m := range rt {
			if !m[nums[i]] {
				m[nums[i]] = true
				res[k] = append(res[k], nums[i])
				flag = true
				break
			}
		}

		if !flag {
			rt = append(rt, map[int]bool{nums[i]: true})
			res = append(res, []int{nums[i]})
		}
	}

	return res
}
