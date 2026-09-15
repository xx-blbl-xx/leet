package main

import (
	"fmt"
	"sort"
)

func threeSum15(nums []int) [][]int {
	sort.Ints(nums)

	k := 1
	for ; k < len(nums)-1; k++ {
		if nums[k] >= 0 {
			break
		}
	}

	i := 0
	j := len(nums) - 1
	res := make([][]int, 0)
	kk := k

	for i < k && k < j {
		m := nums[i] + nums[j]

		if m+nums[kk] == 0 {
			res = append(res, []int{nums[i], nums[kk], nums[j]})
			if kk >= k {
				kk++
			} else {
				kk--
			}
		} else if m > 0 && m+nums[kk] > 0 {
			kk--
		} else if m < 0 && m+nums[kk] < 0 {
			kk++
		} else if m > 0 && m+nums[kk] < 0 {
			kk = j
		} else if m < 0 && m+nums[kk] > 0 {
			kk = i
		} else if m == 0 {
			if kk >= k {
				kk++
			} else {
				kk--
			}
		}

		if kk == i {
			j--
			kk = k
		} else if kk == j {
			i++
			kk = k
		}
	}

	m := make(map[string]bool)
	r := make([][]int, 0)
	for _, v := range res {
		s := fmt.Sprintf("%d_%d_%d", v[0], v[1], v[2])
		if _, ok := m[s]; !ok {
			m[s] = true
			r = append(r, v)
		}
	}

	return r
}
