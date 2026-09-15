package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{strs}
	}

	res := make(map[string][]string, 0)

	for _, str := range strs {
		ha := []rune(str)
		sort.Slice(ha, func(i, j int) bool { return ha[i] > ha[j] })
		hs := string(ha)

		if _, ok := res[hs]; !ok {
			res[hs] = make([]string, 0)
		}
		res[hs] = append(res[hs], str)
	}

	rr := make([][]string, 0, len(res))
	for _, v := range res {
		rr = append(rr, v)
	}

	return rr
}
