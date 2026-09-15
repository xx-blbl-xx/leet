package main

import "sort"

func hIndex(citations []int) int {
	sort.Ints(citations)

	m := 0
	for i := len(citations) - 1; i >= 0 && citations[i] > m; i-- {
		m++
	}

	return m
}
