package dynamic

// generate 杨辉三角 https://leetcode.cn/problems/pascals-triangle/description/?envType=study-plan-v2&envId=top-100-liked
func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	res := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		row[0], row[i] = 1, 1
		for j := 1; j < i; j++ {
			row[j] = res[i-1][j-1] + res[i-1][j]
		}
		res[i] = row
	}
	return res
}
