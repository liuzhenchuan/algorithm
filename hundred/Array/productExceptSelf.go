package Array

// productExceptSelf 除自身以外数组的乘积
func productExceptSelf(nums []int) []int {
	length := len(nums)
	L, R, anwser := make([]int, length), make([]int, length), make([]int, length)

	L[0] = 1
	for i := 1; i < length; i++ {
		L[i] = L[i-1] * nums[i-1]
	}
	R[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		R[i] = R[i+1] * nums[i+1]
	}
	for i := 0; i < length; i++ {
		anwser[i] = L[i] * R[i]
	}
	return anwser
}

func productExceptSelf1(nums []int) []int {
	length := len(nums)
	answer := make([]int, length)

	answer[0] = 1
	for i := 1; i < length; i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}
	R := 1
	for i := length - 1; i >= 0; i-- {
		answer[i] = answer[i] * R
		R = R * nums[i]
	}
	return answer
}
