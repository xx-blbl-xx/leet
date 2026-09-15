package main

func scoreOfString(s string) int {
	res := 0
	for i := 0; i < len(s)-1; i++ {
		t := int(s[i]) - int(s[i+1])
		if t < 0 {
			t = -t
		}
		res += t
	}
	return res
}
