package main

func isValidSudoku(board [][]byte) bool {

	lm := []map[byte]bool{
		{}, {}, {}, {}, {}, {}, {}, {}, {},
	}

	for i, l := range board {
		m := map[byte]bool{}
		for j, v := range l {
			if v == '.' {
				continue
			}
			if m[v] {
				return false
			}
			m[v] = true

			if i < 3 {
				if j < 3 {
					if lm[0][v] {
						return false
					}
					lm[0][v] = true
				} else if j < 6 {
					if lm[1][v] {
						return false
					}
					lm[1][v] = true
				} else {
					if lm[2][v] {
						return false
					}
					lm[2][v] = true
				}
			} else if i < 6 {
				if j < 3 {
					if lm[3][v] {
						return false
					}
					lm[3][v] = true
				} else if j < 6 {
					if lm[4][v] {
						return false
					}
					lm[4][v] = true
				} else {
					if lm[5][v] {
						return false
					}
					lm[5][v] = true
				}
			} else {
				if j < 3 {
					if lm[6][v] {
						return false
					}
					lm[6][v] = true
				} else if j < 6 {
					if lm[7][v] {
						return false
					}
					lm[7][v] = true
				} else {
					if lm[8][v] {
						return false
					}
					lm[8][v] = true
				}
			}

		}
	}

	for i := 0; i < 9; i++ {
		m := map[byte]bool{}
		for j := 0; j < 9; j++ {
			v := board[j][i]
			if v == '.' {
				continue
			}
			if m[v] {
				return false
			}
			m[v] = true
		}
	}

	return true

}
