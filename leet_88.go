package main

import "sort"

func merge2(nums1 []int, m int, nums2 []int, n int) {
	ll := len(nums1)

	for i := n; i < ll; i++ {
		nums1[i] = nums2[i-n]
	}

	sort.Ints(nums1)
}

func merge3(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
		return
	}

	i := m - 1
	j := n - 1
	for {
		if nums2[j] > nums1[i] {
			nums1[i+j+1] = nums2[j]
			j--
		} else {
			nums1[i+j+1] = nums1[i]
			nums1[i] = 0
			i--
		}
		if i < 0 || j < 0 {
			break
		}
	}

	for i := 0; i <= j; i++ {
		nums1[i] = nums2[i]
	}
}
