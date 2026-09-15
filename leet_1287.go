package main

func findSpecialInteger(arr []int) int {
	ll := len(arr)
	num := ll / 4
	mm := make(map[int]int, 0)

	for _, v := range arr {
		n, ok := mm[v]
		if !ok {
			mm[v] = 0
		}
		mm[v]++
		if n >= num {
			return v
		}
	}
	return -1
}
