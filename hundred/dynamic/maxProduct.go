package dynamic

// maxProduct 乘积最大子数组 https://leetcode.cn/problems/maximum-product-subarray/description/?envType=study-plan-v2&envId=top-100-liked
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxProd := make([]int, len(nums))
	minProd := make([]int, len(nums))
	maxProd[0] = nums[0]
	minProd[0] = nums[0]

	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maxProd[i] = max(nums[i], minProd[i-1]*nums[i])
			minProd[i] = min(nums[i], maxProd[i-1]*nums[i])
		} else {
			maxProd[i] = max(nums[i], maxProd[i-1]*nums[i])
			minProd[i] = min(nums[i], minProd[i-1]*nums[i])
		}
		if maxProd[i] > res {
			res = maxProd[i]
		}
	}
	return res
}
