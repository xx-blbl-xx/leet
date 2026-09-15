package main

func countBalls(lowLimit int, highLimit int) int {
	mm := make(map[int]int)
	res := 0
	for i := lowLimit; i <= highLimit; i++ {
		x := getSum(i)
		mm[x]++
		if res < mm[x] {
			res = mm[x]
		}
	}

	return res
}

func getSum(i int) int {
	res := 0
	for i != 0 {
		res += i % 10
		i = i / 10
	}

	return res
}
