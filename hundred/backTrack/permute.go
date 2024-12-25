package backTrack

// permute 全排列 https://leetcode.cn/problems/permutations/description/?envType=study-plan-v2&envId=top-100-liked
func permute(nums []int) [][]int {
	var res [][]int
	var backtrack func(path []int, used []bool)

	// 回溯函数
	backtrack = func(path []int, used []bool) {
		if len(path) == len(nums) { // 结束递归条件
			tmp := make([]int, len(nums))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		// 遍历所有元素
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}

			// 做选择
			used[i] = true
			path = append(path, nums[i])

			// 继续递归
			backtrack(path, used)

			// 回溯，撤销选择
			used[i] = false
			path = path[:len(path)-1]
		}
	}

	// 初始化used数组
	used := make([]bool, len(nums))
	backtrack([]int{}, used)
	return res
}
