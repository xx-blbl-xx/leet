package main

func longestCommonPrefix(strs []string) string {
	res := []byte{}
	l := len(strs[0])
	for i := 0; i < l; i++ {
		ss := strs[0][i]
		flag := true
		for _, s := range strs {
			if len(s)-1 < i || ss != s[i] {
				flag = false
				break
			}
		}
		if flag {
			res = append(res, ss)
		} else {
			break
		}
	}

	return string(res)
}
