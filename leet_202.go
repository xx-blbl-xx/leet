package main

func isHappy(n int) bool {

	mm := make(map[int]bool)
	mm[n] = true
	for {
		sum := 0
		for {
			x := n % 10
			sum += x * x
			y := n / 10
			if y == 0 {
				break
			}
			n = y
		}
		if sum == 1 {
			return true
		}
		if mm[sum] {
			return false
		}
		n = sum
		mm[n] = true
	}

}
