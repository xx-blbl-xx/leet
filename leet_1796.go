package main

func secondHighest(s string) int {
	sb := []byte(s)
	max := byte(0)
	res := byte(0)

	for _, v := range sb {
		if v < 48 || v > 57 {
			continue
		}

		if max < v {
			res = max
			max = v
		} else if max > v && v > res {
			res = v
		}
	}

	if res >= 48 && res != max {
		return int(res) - 48
	}

	return -1
}
