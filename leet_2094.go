package main

import "sort"

func findEvenNumbers(digits []int) []int {

	mm := make(map[int]bool)
	ll := len(digits)
	for i := 0; i < ll; i++ {
		if digits[i] == 0 {
			continue
		}

		for j := 0; j < ll; j++ {
			if j == i {
				continue
			}

			for k := 0; k < ll; k++ {
				if k == i || k == j || digits[k]%2 == 1 {
					continue
				}
				mm[digits[i]*100+digits[j]*10+digits[k]] = true
			}
		}
	}

	res := make([]int, 0, len(mm))
	for k := range mm {
		res = append(res, k)
	}

	sort.Ints(res)

	return res
}
