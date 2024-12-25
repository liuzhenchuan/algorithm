package doublePoint

//盛最多水的容器https://leetcode.cn/problems/container-with-most-water/description/?envType=study-plan-v2&envId=top-100-liked
func maxArea(height []int) int {
	i, j := 0, len(height)-1
	max := 0
	for i < j {
		area := (j - i) * min(height[i], height[j])
		if area > max {
			max = area
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
