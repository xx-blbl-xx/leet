package main

func replaceElements(arr []int) []int {
	ll := len(arr)
	for i := 0; i < ll; i++ {

		max := -1
		for j := i + 1; j < ll; j++ {
			if max < arr[j] {
				max = arr[j]
			}
		}

		arr[i] = max

	}

	return arr
}
