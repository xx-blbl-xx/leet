package main

func trap42(height []int) int {
	if len(height) == 0 {
		return 0
	}

	lm := make([]int, len(height))
	lm[0] = height[0]
	for i := 1; i < len(height); i++ {
		lm[i] = max(height[i], lm[i-1])
	}

	rm := make([]int, len(height))
	rm[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rm[i] = max(height[i], rm[i+1])
	}

	res := 0
	for i := 0; i < len(height); i++ {
		res += min(lm[i], rm[i]) - height[i]
	}

	return res
}
