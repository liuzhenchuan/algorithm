package hash

import (
	"sort"
	"strings"
)

//字母异位词分组https://leetcode.cn/problems/group-anagrams/?envType=study-plan-v2&envId=top-100-liked
// 由于互为字母异位词的两个字符串包含的字母相同，因此对两个字符串分别进行排序之后得到的字符串一定是相同的，故可以将排序之后的字符串作为哈希表的键
func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)
	hashMap := make(map[string][]string)
	for _, str := range strs {
		chars := strings.Split(str, "")
		sort.Strings(chars)
		sortedStr := strings.Join(chars, "")
		hashMap[sortedStr] = append(hashMap[sortedStr], str)
	}
	for _, v := range hashMap {
		result = append(result, v)
	}
	return result
}
