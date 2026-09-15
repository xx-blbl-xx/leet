package main

import "math"

func minSubArrayLen(target int, nums []int) int {
	for _, v := range nums {
		if v >= target {
			return 1
		}
	}

	for m := 2; m <= len(nums); m++ {
		i := 0
		j := m - 1
		rr := 0
		for k := 0; k <= j; k++ {
			rr += nums[k]
		}
		for j < len(nums) {
			if rr >= target {
				return m
			}
			rr -= nums[i]
			i++
			j++
			if j == len(nums) {
				break
			}
			rr += nums[j]
		}

	}

	return 0
}

func minSubArrayLenR(target int, nums []int) int {
	for _, v := range nums {
		if v >= target {
			return 1
		}
	}

	if len(nums) == 1 {
		return 0
	}

	i := 0
	j := 1
	res := 2
	m := nums[i] + nums[j]
	rr := math.MaxInt
	for i < j && j < len(nums) {
		if m >= target {
			rr = min(rr, j-i+1)
			m -= nums[i]
			i++
			res--
		} else {
			j++
			if j == len(nums) {
				if i == 0 && m < target {
					return 0
				}
				break
			}
			m += nums[j]
			res++
		}
	}

	if res > len(nums) {
		return 0
	}

	return rr
}
