package main

func removeDuplicates80(nums []int) int {
	r := 0

	for i := 0; i < len(nums)-2; i++ {
		if r+i+2 == len(nums) {
			break
		}

		if nums[i] == nums[i+1] && nums[i+1] == nums[i+2] {
			r++

			for ii := i + 1; ii < len(nums)-1; ii++ {
				nums[ii] = nums[ii+1]
			}

			if r+i+2 == len(nums) {
				break
			}
			i--
		}

	}

	return len(nums) - r
}
