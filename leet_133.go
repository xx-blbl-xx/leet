package main

import "fmt"

func hammingWeight(num uint32) int {
	ns := fmt.Sprintf("%b", num)

	res := 0
	for _, v := range ns {
		if v == '1' {
			res++
		}
	}

	return res
}
