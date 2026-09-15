package main

import (
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	r := []string{}
	l := 0
	res := []string{}
	for i := 0; i < len(words); {
		k := 0
		k += len(words[i])
		res = []string{words[i]}
		j := 1
		for ; j < 100 && i+j < len(words); j++ {
			k += len(words[i+j]) + 1
			if k <= maxWidth {
				res = append(res, " ")
				res = append(res, words[i+j])
			} else {
				break
			}
		}

		i += j
		l++
		f := fixStr(res, maxWidth)
		r = append(r, f)
	}

	lastStr := strings.Join(res, "")
	ll := len(lastStr)
	for i := 0; i < maxWidth-ll; i++ {
		lastStr += " "
	}
	r[len(r)-1] = lastStr

	return r
}

func fixStr(strs []string, maxWidth int) string {
	w := 0
	sum := 0
	for _, str := range strs {
		if str == " " {
			w++
		}
		sum += len(str)
	}

	if sum == maxWidth {
		return strings.Join(strs, "")
	}

	need := maxWidth - sum
	wi := need
	wj := 0
	if w != 0 {
		wi = need / w
		wj = need % w
	}
	res := []string{}
	for _, str := range strs {
		if str == " " {
			n := wi
			if wj > 0 {
				n++
				wj--
			}
			for i := 0; i < n+1; i++ {
				res = append(res, " ")
			}
		} else {
			res = append(res, str)
		}
	}

	r := strings.Join(res, "")
	ll := len(r)
	for i := 0; i < maxWidth-ll; i++ {
		r += " "
	}

	return r
}
