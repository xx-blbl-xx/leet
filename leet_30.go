package main

import "fmt"

func findSubstring(s string, words []string) []int {
	ll := len(words[0]) * len(words)
	if ll > len(s) {
		return nil
	}

	mm := make([]int, 'z'-'a'+1)
	for _, w := range words {
		for _, v := range w {
			mm[byte(v)-'a']++
		}
	}

	ms := make([]int, 'z'-'a'+1)
	for i := 0; i < ll; i++ {
		ms[s[i]-'a']++
	}

	i := 0
	j := ll - 1
	rr := []int{}
	for j < len(s) {
		r := checkArr(mm, ms)
		if r == 0 {
			fmt.Println(s[i : j+1])
			res := checkWords(s, i, j, words)
			if res != -1 {
				rr = append(rr, res)
				i = j
				j = j + ll
				continue
			}
		}

		k := 0
		for ; k < r; k++ {
			if j+k >= len(s) {
				break
			}
			t := s[i+k]
			t1 := s[j+k]
			ms[t-'a']--
			ms[t1-'a']++
		}

		i += k
		j += k
	}

	return rr
}

func checkArr(mm, ms []int) int {
	r := 0
	for i := 0; i < len(mm); i++ {
		ms[i] = mm[i] - ms[i]
		if ms[i] < 0 {
			r += -ms[i]
		} else {
			r += ms[i]
		}
	}

	return r
}

func checkWords(s string, i, j int, words []string) int {
	l := len(words[0]) - 1

	ww := make([]string, len(words))
	copy(ww, words)
	res := []int{}
	for k := i; k < j; k += l + 1 {
		for t, v := range ww {
			vt := s[k : k+l+1]
			if v == vt {
				ww[t] = ""
				res = append(res, k)
				break
			}
		}

	}

	if len(res) == len(words) {
		return res[0]
	}

	return -1
}
