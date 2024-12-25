package matrix

// setZeroes 矩阵置零
func setZeroes(matrix [][]int) {
	rows := len(matrix)
	if rows == 0 {
		return
	}
	cols := len(matrix[0])

	firstRowHasZero, firstColHasZero := false, false
	// 判断第一行是否有0
	for j := 0; j < cols; j++ {
		if matrix[0][j] == 0 {
			firstRowHasZero = true
		}
	}

	// 判断第一列是否有0
	for i := 1; i < rows; i++ {
		for matrix[i][0] == 0 {
			firstColHasZero = true
		}
	}
	// 遍历非第一行、第一列的数，如果有0，则将对应的第一行、第一列置为0
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
				firstRowHasZero = true
				firstColHasZero = true
			}
		}
	}

	// 遍历非第一行、第一列的数，如果对应的第一行、第一列是0，则将该值置为0
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// 将第一行置0
	if firstRowHasZero {
		for j := 0; j < cols; j++ {
			matrix[0][j] = 0
		}
	}

	// 将第一列置0
	if firstColHasZero {
		for i := 0; i < rows; i++ {
			matrix[i][0] = 0
		}
	}
}
