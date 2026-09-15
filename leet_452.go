package main

import "sort"

func findMinArrowShots(points [][]int) int {
	if len(points) == 1 {
		return 1
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	ma := points[0][1]
	res := 1
	for i := 1; i < len(points); i++ {
		if ma < points[i][0] {
			ma = points[i][1]
			res++
			continue
		} else {
			if ma >= points[i][1] {
				ma = points[i][1]
			}
		}

	}

	return res

}
