package dynamic

// climbStairs 爬楼梯 https://leetcode.cn/problems/climbing-stairs/description/?envType=study-plan-v2&envId=top-100-liked
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	dp := make(map[int]int)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
