package matrix

import "sort"

// searchMatrix 搜索二维矩阵 https://leetcode.cn/problems/search-a-2d-matrix/description/?envType=study-plan-v2&envId=top-100-liked
func searchMatrix(matrix [][]int, target int) bool {
	row := sort.Search(len(matrix), func(i int) bool {
		return matrix[i][0] > target
	}) - 1
	if row < 0 {
		return false
	}
	col := sort.SearchInts(matrix[row], target)
	return col < len(matrix[row]) && matrix[row][col] == target
}
