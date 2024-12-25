package backTrack

// exist 单词搜索 https://leetcode.cn/problems/word-search/description/?envType=study-plan-v2&envId=top-100-liked
func exist(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])

	var backtrack func(r, c, index int) bool
	backtrack = func(r, c, index int) bool {
		if index == len(word) {
			return true
		}

		if r < 0 || r >= rows || c < 0 || c >= cols || board[r][c] != word[index] {
			return false
		}

		temp := board[r][c]
		board[r][c] = '#'

		for _, direction := range directions {
			newRow, newCol := r+direction[0], c+direction[1]
			if backtrack(newRow, newCol, index+1) {
				return true
			}
		}
		board[r][c] = temp
		return false
	}

	// 从每个格子开始尝试
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if board[r][c] == word[0] && backtrack(r, c, 0) {
				return true
			}
		}
	}
	return false
}

var directions = [][2]int{
	{-1, 0}, // 上
	{1, 0},  // 下
	{0, -1}, // 左
	{0, 1},  // 右
}
