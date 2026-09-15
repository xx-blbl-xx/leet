package main

func jump45(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	r := 0
	k := 0
	for i := 0; i < len(nums); {
		m := 0
		for j := 0; j <= nums[i]; j++ {
			if i+j >= len(nums)-1 {
				return r + 1
			}
			if k < j+i+nums[i+j] {
				k = j + i + nums[i+j]
				m = i + j
			}
		}
		i = m

		r++
	}

	return r
}
