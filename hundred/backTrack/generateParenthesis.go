package backTrack

// generateParenthesis 括号生成 https://leetcode.cn/problems/generate-parentheses/description/?envType=study-plan-v2&envId=top-100-liked
func generateParenthesis(n int) []string {
	var res []string
	var path []byte
	var backtrack func(left, right int)
	backtrack = func(left, right int) {
		if left == n && right == n {
			res = append(res, string(path))
			return
		}
		if left < n {
			path = append(path, '(')
			backtrack(left+1, right)
			path = path[:len(path)-1]
		}
		if right < left {
			path = append(path, ')')
			backtrack(left, right+1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0, 0)
	return res
}
