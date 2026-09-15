package main

func removeElement(nums []int, val int) int {
	r := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			r++
			for ii := i; ii < len(nums)-1; ii++ {
				nums[ii] = nums[ii+1]
			}
			i--
			if i+r+1 == len(nums) {
				break
			}
		}
	}

	return len(nums) - r
}
