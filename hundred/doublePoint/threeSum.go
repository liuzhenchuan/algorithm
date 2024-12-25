package doublePoint

import "sort"

// 三数之和 https://leetcode.cn/problems/3sum/description/?envType=study-plan-v2&envId=top-100-liked
func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	result := make([][]int, 0)

	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		value1 := nums[i]
		value2 := 0 - value1

		right := n - 1
		for left := i + 1; left < n; left++ {
			if left > i+1 && nums[left] == nums[left-1] {
				continue
			}
			for left < right && nums[left]+nums[right] > value2 {
				right--
			}
			if left == right {
				break
			}
			if nums[left]+nums[right] == value2 {
				result = append(result, []int{value1, nums[left], nums[right]})
			}
		}
	}
	return result
}
