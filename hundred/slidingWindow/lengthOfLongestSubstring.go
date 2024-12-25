package slidingWindow

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/?envType=study-plan-v2&envId=top-100-liked
// 循环遍历字符串，left记录最长无重复子串开始位置，right记录当前位置
// 双指针，因为要判断字符是否出现过，所以用map存储子串，左指针指向子串头，右指针指向子串尾，左指针初始化为第0位，然后不断遍历到最后一位
// 右指针初始化未-1，向右移动一位，边界是下一位是最后一位，如果下一位未出现过，继续右移一位，子串插入一位
// 如果下一位出现过，则从左指针开始删除，左指针右移，直到右指针出现的下一位
func lengthOfLongestSubstring(s string) int {
	right := -1
	max := 0
	l := len(s)
	m := make(map[byte]bool)
	for i := 0; i < l; i++ {
		for right+1 < l && !m[s[right+1]] {
			m[s[right+1]] = true
			right++
		}
		if len(m) > max {
			max = len(m)
		}
		delete(m, s[i])
	}
	return max
}
