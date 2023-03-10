func isValidSudoku(board [][]byte) bool {
	col := make(map[int][]string)
	row := make(map[int][]string)
	sq := make(map[int][]string)

	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			temp := string(board[x][y])
			if temp == "." {
				continue
			}

			if val, exists := col[x]; exists && containsElement(val, temp) {
                    return false
            }

            if val, exists := row[y]; exists && containsElement(val, temp) {
                    return false
            }

			if val, exists := sq[x/3+y/3*10]; exists && containsElement(val, temp) {
                    return false
            }
			col[x] = append(col[x], temp)
			row[y] = append(row[y], temp)
			sq[x/3+y/3*10] = append(sq[x/3+y/3*10], temp)
		}
	}

	return true
}

func containsElement(s []string, v string) bool {
    for _, i := range(s) {
        if v == i {
            return true
        }
    }
    return false
}


