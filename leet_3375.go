package main

import "sort"

func minOperations(nums []int, k int) int {
	sort.Ints(nums)
	if nums[0] < k {
		return -1
	}

	res := 0
	num := k
	for _, v := range nums {
		if num != v {
			res++
			num = v
		}
	}
	return res
}
