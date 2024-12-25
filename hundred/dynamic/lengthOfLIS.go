package dynamic

// lengthOfLIS 最长递增子序列 https://leetcode.cn/problems/longest-increasing-subsequence/description/?envType=study-plan-v2&envId=top-100-liked
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	maxLIS := 0
	for _, length := range dp {
		maxLIS = max(maxLIS, length)
	}
	return maxLIS
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
