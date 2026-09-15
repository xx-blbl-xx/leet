package main

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int][]int)

	for k, v := range nums {
		if _, ok := m[v]; !ok {
			m[v] = []int{}
		}
		m[v] = append(m[v], k)
	}

	for _, ll := range m {
		if len(ll) == 1 {
			continue
		}

		for i := 0; i < len(ll)-1; i++ {
			if ll[i+1]-ll[i] > k {
				continue
			} else {
				return true
			}
		}
	}

	return false
}
