package hash

// 两数之和 https://leetcode.cn/problems/two-sum/submissions/512577899/?envType=study-plan-v2&envId=top-100-liked

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	result := make([]int, 2)
	for k, v := range nums {
		m[v] = k
	}
	for k, v := range nums {
		value := target - v
		if _, ok := m[value]; ok && m[value] != k {
			result[0] = k
			result[1] = m[value]
		}
	}
	return result
}
