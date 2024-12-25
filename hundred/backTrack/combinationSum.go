package backTrack

// combinationSum 组合总和 https://leetcode.cn/problems/combination-sum/description/?envType=study-plan-v2&envId=top-100-liked
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var path []int

	var backtrack func(start int, target int)
	backtrack = func(start int, target int) {
		if target == 0 {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := start; i < len(candidates); i++ {
			if candidates[i] > target {
				continue
			}
			path = append(path, candidates[i])
			backtrack(i, target-candidates[i])
			path = path[:len(path)-1]
		}
	}
	backtrack(0, target)
	return res
}
