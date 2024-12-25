package binarySearch

// searchInsert 搜索插入位置 https://leetcode.cn/problems/search-insert-position/description/?envType=study-plan-v2&envId=top-100-liked
func searchInsert(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
