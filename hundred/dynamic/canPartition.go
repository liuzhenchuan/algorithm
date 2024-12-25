package dynamic

// canPartition 分割等和子集 https://leetcode.cn/problems/partition-equal-subset-sum/description/?envType=study-plan-v2&envId=top-100-liked
func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	dp := make([]bool, target+1)
	dp[0] = true

	for _, num := range nums {
		for i := target; i >= num; i-- {
			if dp[i-num] {
				dp[i] = true
			}
		}
	}
	return dp[target]
}
