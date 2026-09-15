package main

import "strconv"

func divisorSubstrings(num int, k int) int {
	s := strconv.Itoa(num)
	n := len(s)
	res := 0
	for i := 0; i <= n-k; i++ {
		t, _ := strconv.Atoi(s[i : i+k])
		if t != 0 && num%t == 0 {
			res++
		}
	}
	return res
}
