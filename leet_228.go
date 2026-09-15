package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	res := []string{}
	for i := 0; i < len(nums); {

		j := i + 1
		for ; j < len(nums); j++ {
			if nums[j]-nums[j-1] != 1 {
				break
			}
		}

		if j == i+1 {
			res = append(res, strconv.FormatInt(int64(nums[i]), 10))
			i++
		} else {
			res = append(res, fmt.Sprintf("%d->%d", nums[i], nums[j-1]))
			i = j
		}

	}

	return res
}
