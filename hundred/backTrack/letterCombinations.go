package backTrack

// letterCombinations 电话号码的字母组合 https://leetcode.cn/problems/letter-combinations-of-a-phone-number/description/?envType=study-plan-v2&envId=top-100-liked
func letterCombinations(digits string) []string {
	var res []string
	if len(digits) == 0 {
		return res
	}
	var backtrack func(index int, path string)
	backtrack = func(index int, path string) {
		if index == len(digits) {
			res = append(res, path)
			return
		}
		letters := digitToLetters[rune(digits[index])]
		for i := 0; i < len(letters); i++ {
			backtrack(index+1, path+string(letters[i]))
		}
	}
	backtrack(0, "")
	return res
}

var digitToLetters = map[rune]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}
