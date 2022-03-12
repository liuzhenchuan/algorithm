// 最长公共子串 https://blog.nowcoder.net/n/05ea4d923b544c96bc07f4a4a8437331
package string

func LCS(str1, str2 string) string {
	len1 := len(str1)
	len2 := len(str2)
	if len1 == 0 || len2 == 0 {
		return ""
	}

	dp := make([][]int, len1+1)
	for i := 0; i < len1+1; i++ {
		dp[i] = make([]int, len2+1)
	}

	max := 0
	end := 0

	for i := 1; i < len1+1; i++ {
		for j := 1; j < len2+1; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = 0
			}
			if dp[i][j] > max {
				max = dp[i][j]
				end = i
			}
		}
	}

	if max == 0 {
		return ""
	}
	return str1[end-max : end]
}

func LCS1(str1 string, str2 string) string {
	len1 := len(str1)
	len2 := len(str2)
	if len1 == 0 || len2 == 0 {
		return ""
	}

	dp := make([][]int, len1+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len2+1)
	}

	ans := 0
	h := map[int]int{}
	for i := len1 - 1; i >= 0; i-- {
		for j := len2 - 1; j >= 0; j-- {
			if str1[i] == str2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				dp[i][j] = 0
			}
			if dp[i][j] > ans {
				ans = dp[i][j]
				h[ans] = i
			}
		}
	}

	start := h[ans]
	return str1[start : start+ans]
}

// 最长公共子序列
func LCSS(text1 string, text2 string) int {
	len1, len2 := len(text1), len(text2)
	if len1 == 0 || len2 == 0 {
		return 0
	}
	dp := make([][]int, len1+1)
	for i := range dp {
		dp[i] = make([]int, len2+1)
	}
	for i, c1 := range text1 {
		for j, c2 := range text2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[len1][len2]
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
