package main

func removeDuplicates(nums []int) int {
	r := 0
	last := -1000
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == -10000 {
			break
		}
		if nums[i] == nums[i+1] {
			if last != nums[i] {
				last = nums[i]
			}
			r++
			for ii := i; ii < len(nums)-1; ii++ {
				nums[ii] = nums[ii+1]
			}
			nums[len(nums)-1] = -10000
			i--
			if 1+i+r == len(nums) {
				break
			}
		}
	}

	return len(nums) - r
}
