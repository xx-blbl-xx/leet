package main

func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	m2 := map[byte][]byte{
		'I': []byte{'V', 'X'},
		'X': []byte{'L', 'C'},
		'C': []byte{'D', 'M'},
	}

	res := 0
	for i := 0; i < len(s); i++ {
		f := 1
		if i < len(s)-1 {
			arr := m2[s[i]]
			if len(arr) != 0 && (arr[0] == s[i+1] || arr[1] == s[i+1]) {
				f = -1
			}
		}
		res += f * m[s[i]]
	}

	return res
}
