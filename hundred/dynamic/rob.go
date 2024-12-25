package dynamic

// rob 打家劫舍 https://leetcode.cn/problems/house-robber/description/?envType=study-plan-v2&envId=top-100-liked
func rob(nums []int) int {
	// dp[i] = max(nums[i] + dp[i-2], dp[i-1])
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		} else {
			return nums[1]
		}
	}
	dp := make(map[int]int)
	dp[1] = nums[0]
	if nums[0] > nums[1] {
		dp[2] = nums[0]
	} else {
		dp[2] = nums[1]
	}
	for i := 3; i <= len(nums); i++ {
		if dp[i-2]+nums[i-1] > dp[i-1] {
			dp[i] = dp[i-2] + nums[i-1]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(nums)]
}
