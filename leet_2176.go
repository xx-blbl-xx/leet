package main

func countPairs(nums []int, k int) int {
	mm := make(map[int][]int)

	for k, v := range nums {
		if len(mm[v]) == 0 {
			mm[v] = make([]int, 0)
		}
		mm[v] = append(mm[v], k)
	}

	res := 0
	for _, ll := range mm {
		if len(ll) < 2 {
			continue
		}

		for i := 0; i < len(ll)-1; i++ {
			for j := i + 1; j < len(ll); j++ {
				if (ll[i]*ll[j])%k == 0 {
					res++
				}
			}
		}
	}

	return res
}
