package main

func canCompleteCircuit(gas []int, cost []int) int {
	if len(gas) == 1 && gas[0] >= cost[0] {
		return 0
	}
	for i := 0; i < len(gas); i++ {
		if gas[i]-cost[i] <= 0 {
			continue
		}

		g := 0
		c := 0
		j := 0
		k := 0
		for ; j < len(gas); j++ {
			k = (i + j) % len(gas)
			g += gas[k]
			c += cost[k]

			if g-c < 0 {
				break
			}

		}
		if j == len(gas) && g-c >= 0 {
			return i
		}
	}

	return -1
}
