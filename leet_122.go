package main

func maxProfit122(prices []int) int {
	r := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] {
			n := prices[i+1] - prices[i]
			r += max(0, n)
		}
	}

	return r
}
