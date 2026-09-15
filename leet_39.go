package main

import (
	"fmt"
	"sort"
)

func combinationSum39(candidates []int, target int) [][]int {
	type obj struct {
		nums []int
		sum  int
	}
	res := []*obj{}
	can := []int{}
	rr := map[string][]int{}

	for _, v := range candidates {
		if v > target {
			continue
		}
		if v == target {
			rr[""] = []int{v}
			continue
		}
		res = append(res, &obj{
			nums: []int{v}, sum: v,
		})
		can = append(can, v)
	}

	f := func(rs []*obj, ns []int) []*obj {
		nrs := []*obj{}
		for _, v := range ns {
			for _, r := range rs {
				nr := &obj{
					nums: append([]int{}, r.nums...),
					sum:  r.sum,
				}
				nr.nums = append(nr.nums, v)
				nr.sum += v
				if nr.sum > target {
					continue
				}
				if nr.sum == target {
					sort.Ints(nr.nums)
					str := fmt.Sprintf("%+v", nr.nums)
					rr[str] = nr.nums
					continue
				}
				nrs = append(nrs, nr)
			}
		}
		return nrs
	}
	rs := res
	for i := 0; i < 40; i++ {
		rs = f(rs, can)
		if len(rs) == 0 {
			break
		}
	}

	rrs := [][]int{}
	for _, s := range rr {
		rrs = append(rrs, s)
	}

	return rrs
}
