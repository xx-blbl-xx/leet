package main

func intToRoman(num int) string {

	res := []byte{}

	if num >= 1000 {
		n := num / 1000
		for range n {
			res = append(res, 'M')
		}
		num -= n * 1000
	}

	if num >= 100 {
		n := num / 100
		if n == 9 {
			res = append(res, 'C', 'M')
			num -= 900
		} else if n == 4 {
			res = append(res, 'C', 'D')
			num -= 400
		} else if n >= 5 {
			res = append(res, 'D')
			for range n - 5 {
				res = append(res, 'C')
			}
			num -= n * 100
		} else {
			for range n {
				res = append(res, 'C')
			}
			num -= n * 100
		}
	}

	if num >= 10 {
		n := num / 10
		if n == 9 {
			res = append(res, 'X', 'C')
			num -= 90
		} else if n == 4 {
			res = append(res, 'X', 'L')
			num -= 40
		} else if n >= 5 {
			res = append(res, 'L')
			for range n - 5 {
				res = append(res, 'X')
			}
			num -= n * 10
		} else {
			for range n {
				res = append(res, 'X')
			}
			num -= n * 10
		}
	}

	n := num
	if n == 9 {
		res = append(res, 'I', 'X')
		num -= 9
	} else if n == 4 {
		res = append(res, 'I', 'V')
		num -= 4
	} else if n >= 5 {
		res = append(res, 'V')
		for range n - 5 {
			res = append(res, 'I')
		}
		num -= n * 1
	} else {
		for range n {
			res = append(res, 'I')
		}
		num -= n * 1
	}

	return string(res)
}
