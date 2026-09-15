package main

import "fmt"

func twoSum1672(numbers []int, target int) []int {
	i := 0
	j := 1
	for i < len(numbers) && j < len(numbers) {
		fmt.Println(numbers[i]+numbers[j], i, j)
		if numbers[i]+numbers[j] == target {
			break
		}
		if numbers[i]+numbers[j] < target {
			if i < j {
				i, j = j, i
			}
			if j == len(numbers)-1 {
				i++
				j = i + 1
			} else {
				j++
			}
		}

		if numbers[i]+numbers[j] > target {
			if i < j {
				j--
				i++
				i, j = j, i
			} else {
				i--
				j++
			}
		}
	}

	return []int{min(i, j) + 1, max(i, j) + 1}
}

func twoSum167(numbers []int, target int) []int {
	i := 0
	j := 1
	for i < len(numbers) && j < len(numbers) {
		if numbers[i]+numbers[j] == target {
			break
		}
		if numbers[i]+numbers[j] < target {
			if j == len(numbers)-1 {
				i++
				j = i + 1
			} else {
				j++
			}
		}

		if numbers[i]+numbers[j] > target {
			i++
			j = i + 1
		}
	}

	return []int{min(i, j) + 1, max(i, j) + 1}
}
