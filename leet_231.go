package main

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}

	i := 1
	for n/(1<<i) != 0 {
		x := n % (1 << i)
		if x != 0 {
			return false
		}
		i++
	}

	return true
}
