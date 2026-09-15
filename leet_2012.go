package main

func sumOfBeauties(nums []int) int {
	res := 0
	ll := len(nums)

	n2 := make([]int, ll)
	pm := nums[0]
	for i := 1; i < ll-1; i++ {
		if pm < nums[i] {
			pm = nums[i]
			n2[i] = 1
		}
	}

	lm := nums[ll-1]
	for i := ll - 2; i > 0; i-- {
		if n2[i] == 1 && lm > nums[i] {
			res += 2
		} else if nums[i-1] < nums[i] && nums[i] < nums[i+1] {
			res++
		}

		lm = min(lm, nums[i])
	}

	return res
}
