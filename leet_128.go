package main

func longestConsecutive128(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	m := map[int]int{}

	for _, v := range nums {
		m[v] = 1
	}

	res := 1
	for k := range m {
		i := k

		for {
			i++
			if _, ok := m[i]; !ok {
				break
			}
			m[k] += m[i]
			delete(m, i)
		}
		res = max(res, m[k])

	}

	return res

}
