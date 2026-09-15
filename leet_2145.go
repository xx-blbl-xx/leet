package main

func numberOfArrays(differences []int, lower int, upper int) int {
	cur := 0
	maxX := lower
	miny := upper
	for _, d := range differences {

		cur += d
		x := lower + cur
		y := upper + cur

		maxX = max(maxX, x)
		miny = min(miny, y)
		if maxX > upper || miny < lower {
			return 0
		}
	}
	if miny < maxX {
		return 0
	}
	return miny - maxX + 1
}
