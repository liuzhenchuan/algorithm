package doublePoint

// 移动零 https://leetcode.cn/problems/move-zeroes/description/?envType=study-plan-v2&envId=top-100-liked
// 左指针左面全是非零数，右指针右面全是待排的，如果有非0数，交换左右指针的数
func moveZeroes(nums []int) {
	i, j, n := 0, 0, len(nums)
	for j < n {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		j++
	}
}
