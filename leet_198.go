package main

func rob198(nums []int) int {
	l := len(nums)

	switch l {
	case 1:
		return nums[0]
	case 2:
		return max(nums[0], nums[1])
	case 3:
		return max(nums[0]+nums[2], nums[1])
	}

	res := make([]int, l)

	res[0] = nums[0]
	res[1] = max(nums[0], nums[1])
	res[2] = max(nums[0]+nums[2], nums[1])

	i := 3
	for ; i < l; i++ {
		res[i] = nums[i] + max(res[i-3], res[i-2])
	}

	rr := max(res[i-1], res[i-2])

	return rr
}
