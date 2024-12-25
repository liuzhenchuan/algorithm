package Array

// maxSubArray 最大子数组和
func maxSubArray(nums []int) int {
	// 初始化最大值
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		// 如果后一个值加上之前的和大于自己，则自己累加， 如果小于自己，则自己是最大值
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		// 当前最大值和历史最大值进行比较
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
