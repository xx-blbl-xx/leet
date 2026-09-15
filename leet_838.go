package main

func pushDominoes(dominoes string) string {
	ds := []byte(dominoes)
	r := -1
	if ds[0] == 'R' {
		r = 0
	}
	for i := 1; i < len(ds); i++ {
		if ds[i] == '.' {
			if ds[i-1] == 'R' {
				ds[i] = 'R'
			}
			continue
		}
		if ds[i] == 'R' {
			r = i
			continue
		}
		if ds[i] == 'L' {
			_, end := setLeft(ds, r, i)
			r = end
		}
	}
	return string(ds)
}

func setLeft(ds []byte, start, end int) (int, int) {
	if start < 0 {
		start = 0
	} else if ds[start] == 'L' {

	} else {
		mid := (start + end)
		if mid%2 == 0 {
			ds[mid/2] = '.'
		}
		start = mid/2 + 1
	}

	for i := start; i <= end; i++ {
		ds[i] = 'L'
	}

	return start, end
}
