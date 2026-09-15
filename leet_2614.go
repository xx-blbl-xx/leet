package main

import "sort"

func diagonalPrime(nums [][]int) int {
	ll := len(nums)
	res := 0
	resArr := []int{}
	for i := 0; i < ll; i++ {
		resArr = append(resArr, nums[i][i], nums[i][ll-i-1])
	}
	sort.Ints(resArr)
	for i := len(resArr) - 1; i >= 0; i-- {
		if isPrime(resArr[i]) {
			res = resArr[i]
			break
		}
	}

	return res
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
