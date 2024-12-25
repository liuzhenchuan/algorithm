package subString

// 和为k的子数组  https://leetcode.cn/problems/subarray-sum-equals-k/description/?envType=study-plan-v2&envId=top-100-liked
// 以i为结尾的子串和为k，遍历
func subarraySum1(nums []int, k int) int {
	l := len(nums)
	ans := 0
	if l == 0 {
		return 0
	}
	for i := 0; i < l; i++ {
		sum := 0
		for end := i; end >= 0; end-- {
			sum = sum + nums[end]
			if sum == k {
				ans++
			}
		}
	}
	return ans
}

// pre[i] 表示i之前的所有和， [j,i] 和为k, 那么pre[j-1] = pre[i] - k,累计和，如果符合这个条件计1
func subarraySum2(nums []int, k int) int {
	l := len(nums)
	ans := 0
	total := 0
	if l == 0 {
		return 0
	}
	m := make(map[int]int) // key为pre和
	m[0] = 1
	for i := 0; i < l; i++ {
		total += nums[i]
		if _, ok := m[total-k]; ok {
			ans += m[total-k]
		}
		m[total]++
	}
	return ans
}
