package stack

func decodeString(s string) string {
	stack := []rune{}
	num := 0
	result := ""

	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			num = num*10 + int(ch-'0')
		} else if ch == '[' {
			stack = append(stack, rune(num))
			num = 0
			stack = append(stack, ch)
		} else if ch == ']' {
			temp := ""
			for len(stack) > 0 && stack[len(stack)-1] != '[' {
				temp = string(stack[len(stack)-1]) + temp
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			repeatCount := int(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			repeated := ""
			for i := 0; i < repeatCount; i++ {
				repeated += temp
			}
			stack = append(stack, []rune(repeated)...)
		} else {
			stack = append(stack, ch)
		}
	}
	for _, ch := range stack {
		result += string(ch)
	}
	return result
}
