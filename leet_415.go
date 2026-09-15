package main

func addStrings(num1 string, num2 string) string {
	b1 := []byte(num1)
	b2 := []byte(num2)
	l1 := len(b1)
	l2 := len(b2)

	ma := max(l1, l2)

	res := make([]byte, ma)

	f := byte(0)
	for i := 0; i < ma; i++ {
		p := l1 - 1 - i
		q := l2 - 1 - i
		b1n := byte(0)
		b2n := byte(0)
		if p >= 0 {
			b1n = b1[p]
		}
		if q >= 0 {
			b2n = b2[q]
		}
		x := b1n + b2n + f
		if x >= 96 {
			x -= 48
		}

		if x > 57 {
			x -= 10
			f = 1
		} else {
			f = 0
		}

		res[ma-1-i] = x
	}
	if f == 1 {
		return "1" + string(res)
	}

	return string(res)
}
