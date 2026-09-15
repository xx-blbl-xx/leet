package main

import "fmt"

func evenOddBit(n int) []int {
	str := fmt.Sprintf("%b", n)

	res := []int{0, 0}
	ll := len(str)
	for i := ll - 1; i >= 0; i-- {
		if str[i] == '1' {
			if (ll-i-1)%2 == 0 {
				res[0]++
			} else {
				res[1]++
			}
		}
	}

	return res
}
