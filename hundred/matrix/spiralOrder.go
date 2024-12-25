package matrix

// spiralOrder 螺旋矩阵
func spiralOrder(matrix [][]int) []int {
	// 为空判断
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	// 行数，列数
	rows, columns := len(matrix), len(matrix[0])
	// 初始化已访问过的矩阵
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, columns)
	}

	var (
		total          = rows * columns                                                // 总数
		order          = make([]int, total)                                            // 返回结果
		row, column    = 0, 0                                                          // 横纵坐标
		directions     = [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}} //定义方向
		directionIndex = 0
	)
	for i := 0; i < total; i++ {
		order[i] = matrix[row][column]
		visited[row][column] = true
		nextRow, nextColumn := row+directions[directionIndex][0], column+directions[directionIndex][1]
		if nextRow < 0 || nextRow >= rows || nextColumn < 0 || nextColumn >= columns || visited[nextRow][nextColumn] {
			directionIndex = (directionIndex + 1) % 4
		}
		row = nextRow
		column = nextColumn
	}
	return order
}
