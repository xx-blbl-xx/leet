package main

func threeConsecutiveOdds(arr []int) bool {
	res := 0

	for _, v := range arr {
		if v%2 == 0 {
			res = 0
			continue
		}
		res++
		if res >= 3 {
			return true
		}
	}

	return false
}
