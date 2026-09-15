package main

func plusOne(digits []int) []int {
	l := len(digits)
	for i := l - 1; i >= 0; i-- {
		r := digits[i]
		if r != 9 {
			digits[i] += 1
			break
		}
		digits[i] = 0
		if i == 0 {
			digits = append([]int{1}, digits...)
			break
		}
	}

	return digits
}
