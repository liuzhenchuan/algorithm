package stack

// isValid 有效的括号 https://leetcode.cn/problems/valid-parentheses/description/?envType=study-plan-v2&envId=top-100-liked
func isValid(s string) bool {
	stack := []rune{}
	pair := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, ch := range s {
		if opening, ok := pair[ch]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != opening {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}
	return len(stack) == 0
}
