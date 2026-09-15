package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	l := make([]int, 0)
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+":
			j := l[len(l)-2]
			k := l[len(l)-1]
			l = l[:len(l)-2]
			l = append(l, j+k)
		case "-":
			j := l[len(l)-2]
			k := l[len(l)-1]
			l = l[:len(l)-2]
			l = append(l, j-k)
		case "*":
			j := l[len(l)-2]
			k := l[len(l)-1]
			l = l[:len(l)-2]
			l = append(l, j*k)
		case "/":
			j := l[len(l)-2]
			k := l[len(l)-1]
			l = l[:len(l)-2]
			l = append(l, j/k)
		default:
			num, err := strconv.ParseInt(tokens[i], 10, 64)
			if err != nil {
				fmt.Println(err)
				return 0
			}
			l = append(l, int(num))
		}
	}

	return l[0]
}
