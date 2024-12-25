package binarySearch

// search 搜索旋转排序数组 https://leetcode.cn/problems/search-in-rotated-sorted-array/description/?envType=study-plan-v2&envId=top-100-liked
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] { // 左边有序部分
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
