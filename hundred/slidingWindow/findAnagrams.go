package slidingWindow

// 找到字符串中所有字母异位词 https://leetcode.cn/problems/find-all-anagrams-in-a-string/description/?envType=study-plan-v2&envId=top-100-liked
// 给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
//
//异位词 指由相同字母重排列形成的字符串（包括相同的字符串）
// 滑动窗口思想， right右移扩大窗口，删除前面，left右移缩小窗口

// 找到字符串中所有字母异位词
// 通过滑动窗口，从左右到右，判断每个子串是否是字母异位词
// 针对go语言，采用24位的数组相等来判断是否是字母异位词， 每次窗口滑动时，增加right字母的次数，并减少left字母次数
func findAnagrams(s string, p string) []int {
	sl := len(s)
	pl := len(p)
	ans := make([]int, 0)
	if sl < pl {
		return ans
	}

	var sArr, pArr [26]int
	for k, v := range p {
		sArr[s[k]-'a']++
		pArr[v-'a']++
	}
	if sArr == pArr {
		ans = append(ans, 0)
	}
	for k, v := range s[:sl-pl] {
		sArr[v-'a']--
		sArr[s[k+pl]-'a']++
		if sArr == pArr {
			ans = append(ans, k+1)
		}
	}
	return ans
}
