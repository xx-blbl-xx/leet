package main

func isBalanced(num string) bool {
	r1 := 0
	r2 := 0
	for i := range num {
		if i%2 == 0 {
			r1 += int(num[i] - '0')
		} else {
			r2 += int(num[i] - '0')
		}
	}

	return r1 == r2
}
