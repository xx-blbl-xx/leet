package main

func canJump55(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	k := 0
	for i := 0; i < len(nums); i++ {
		if i > k {
			return false
		}

		k = max(nums[i]+i, k)
	}

	return true
}
