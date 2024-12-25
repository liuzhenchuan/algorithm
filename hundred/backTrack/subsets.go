package backTrack

// subsets 子集 https://leetcode.cn/problems/subsets/description/?envType=study-plan-v2&envId=top-100-liked
func subsets(nums []int) [][]int {
	var res [][]int
	var path []int
	var backtrack func(start int)
	backtrack = func(start int) {
		res = append(res, append([]int{}, path...))

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return res
}
