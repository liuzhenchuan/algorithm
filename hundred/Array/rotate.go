package Array

// Rotate 旋转数组
func Rotate(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func Rotate1(nums []int, k int) {
	k = k % len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}
