package main

func maxArea11(height []int) int {
	res := 0
	i, j := 0, len(height)-1
	mj := 0
	for i <= j {
		if height[j] < mj {
			if j == i {
				j = len(height) - 1
				i++
				mj = 0
			} else {
				j--
			}
			continue
		}
		mj = height[j]
		m := (j - i) * min(height[i], height[j])
		if m > res {
			res = m
		}

		if j == i {
			j = len(height) - 1
			i++
			mj = 0
		} else {
			j--
		}

	}

	return res
}

// 总是移动最短的那个
func maxArea11R(height []int) int {
	res := 0
	i, j := 0, len(height)-1

	for i <= j {
		m := (j - i) * min(height[i], height[j])
		if m > res {
			res = m
		}

		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}

	return res
}
