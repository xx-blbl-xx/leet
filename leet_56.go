package main

import "sort"

func merge56(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0)

	mi := intervals[0][0]
	ma := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if ma < intervals[i][0] {
			res = append(res, []int{mi, ma})
			mi = intervals[i][0]
			ma = intervals[i][1]
		} else {
			if ma <= intervals[i][1] {
				ma = intervals[i][1]
			}
		}
	}

	res = append(res, []int{mi, ma})
	return res
}
