package main

func lengthOfLIS300(nums []int) int {
	ll := len(nums)

	if ll == 1 {
		return 1
	}

	res := make([]int, ll)
	res[0] = 1
	if nums[0] < nums[1] {
		res[1] = 2
	} else {
		res[1] = 1
	}

	maxRes := max(res[0], res[1])
	for i := 2; i < ll; i++ {

		res[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				res[i] = max(res[j]+1, res[i])
			}
		}

		maxRes = max(maxRes, res[i])
	}

	return maxRes

}
