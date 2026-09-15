package main

func minDominoRotations(tops []int, bottoms []int) int {
	res1 := checkSame(tops[0], tops, bottoms)
	res2 := checkSame(tops[0], bottoms, tops)
	res3 := checkSame(bottoms[0], bottoms, tops)
	res4 := checkSame(bottoms[0], tops, bottoms)

	r1 := 0
	if res1 == -1 && res2 == -1 {
		r1 = -1
	} else if res1 == -1 {
		r1 = res2
	} else if res2 == -1 {
		r1 = res1
	} else {
		r1 = min(res1, res2)
	}

	r2 := 0
	if res3 == -1 && res4 == -1 {
		r2 = -1
	} else if res3 == -1 {
		r2 = res4
	} else if res4 == -1 {
		r2 = res3
	} else {
		r2 = min(res3, res4)
	}

	if r1 == -1 && r2 == -1 {
		return -1
	} else if r1 == -1 {
		return r2
	} else if r2 == -1 {
		return r1
	}

	return min(r1, r2)
}

func checkSame(tmp int, tops []int, bottoms []int) int {
	res := 0
	for i := 0; i < len(tops); i++ {
		if tops[i] == tmp {
			continue
		}
		if bottoms[i] == tmp {
			res++
			continue
		}
		return -1
	}

	return min(res, len(tops)-res)
}
