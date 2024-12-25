package dynamic

// numSquares 完全平方数 https://leetcode.cn/problems/perfect-squares/description/?envType=study-plan-v2&envId=top-100-liked
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i
	}

	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
