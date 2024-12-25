package backTrack

// partition 分割回文串 https://leetcode.cn/problems/palindrome-partitioning/description/?envType=study-plan-v2&envId=top-100-liked
func partition(s string) [][]string {
	var res [][]string
	backtrack(s, 0, []string{}, &res)
	return res
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func backtrack(s string, start int, path []string, res *[][]string) {
	if start == len(s) {
		*res = append(*res, append([]string{}, path...))
		return
	}
	for end := start + 1; end <= len(s); end++ {
		if isPalindrome(s[start:end]) {
			path = append(path, s[start:end])
			backtrack(s, end, path, res)
			path = path[:len(path)-1]
		}
	}
}
