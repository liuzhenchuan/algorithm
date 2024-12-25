package binarySearch

// findMin 寻找旋转数组最小值 https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/description/?envType=study-plan-v2&envId=top-100-liked
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	if nums[left] <= nums[right] {
		return nums[left]
	}

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]

}
