package main

import "sort"

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	res := append(intervals, newInterval)
	sort.Slice(res, func(i, j int) bool {
		return res[i][0] < res[j][0]
	})

	mi := res[0][0]
	ma := res[0][1]
	rr := make([][]int, 0)
	for i := 1; i < len(res); i++ {
		if ma < res[i][0] {
			rr = append(rr, []int{mi, ma})
			mi = res[i][0]
			ma = res[i][1]
			continue
		} else {
			if ma < res[i][1] {
				ma = res[i][1]
			}
		}

	}

	rr = append(rr, []int{mi, ma})

	return rr
}
