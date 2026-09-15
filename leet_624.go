package main

func maxDistance(arrays [][]int) int {
	s := 200
	sk := -1
	s2 := 201
	b := -200
	bk := -1
	b2 := -199

	for i := 0; i < len(arrays); i++ {
		ll := len(arrays[i])
		if arrays[i][0] <= s2 {
			if arrays[i][0] <= s {
				s2 = s
				s = arrays[i][0]
				sk = i
			} else {
				s2 = arrays[i][0]
			}
		}

		if arrays[i][ll-1] >= b2 {
			if arrays[i][ll-1] >= b {
				b2 = b
				b = arrays[i][ll-1]
				bk = i
			} else {
				b2 = arrays[i][ll-1]
			}
		}
	}

	if sk != bk {
		res := b - s
		return max(res, -res)
	}
	return max(b-s2, s2-b, b2-s, s-b2)
}
