package main

func coinChange322(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	res := make([]int, amount+1)

	i := coins[0]
	for _, v := range coins {
		i = min(v, i)
	}

	for ; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] == i {
				res[i] = 1
				break
			}

			if i-coins[j] > 0 && res[i-coins[j]] != 0 {
				if res[i] == 0 {
					res[i] = res[i-coins[j]] + 1
				} else {
					res[i] = min(res[i], res[i-coins[j]]+1)
				}
			}
		}

	}

	if res[amount] == 0 {
		return -1
	}

	return res[amount]
}
