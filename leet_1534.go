package main

func countGoodTriplets(arr []int, a int, b int, c int) int {
	ll := len(arr)
	res := 0
	for i := 0; i < ll-2; i++ {
		for j := i + 1; j < ll-1; j++ {
			for k := j + 1; k < ll; k++ {
				if abs(arr[i]-arr[j]) <= a && abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c {
					res++
				}
			}
		}
	}

	return res
}
