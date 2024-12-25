package stack

// dailyTemperatures 每日温度 https://leetcode.cn/problems/daily-temperatures/description/?envType=study-plan-v2&envId=top-100-liked
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	result := make([]int, n)
	stack := []int{}

	for i := 0; i < n; i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[index] = i - index
		}
		stack = append(stack, i)
	}
	return result
}
