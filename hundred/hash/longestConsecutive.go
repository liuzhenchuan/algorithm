package hash

// 最长连续序列 https://leetcode.cn/problems/longest-consecutive-sequence/description/?envType=study-plan-v2&envId=top-100-liked

func longestConsecutive(nums []int) int {
	// 通过但超时
	//hashMap := make(map[int]bool)
	//max := 0
	//for _, num := range nums {
	//	hashMap[num] = true
	//}
	//
	//for _, num := range nums {
	//	if _, ok := hashMap[num-1]; ok {
	//		continue
	//	}
	//	l := 1
	//	for {
	//		if _, ok := hashMap[num+l]; ok {
	//			l++
	//		} else {
	//			break
	//		}
	//	}
	//	if l > max {
	//		max = l
	//	}
	//}
	//return max

	hashMap := make(map[int]bool)
	for _, num := range nums {
		hashMap[num] = true
	}
	max := 0
	for num := range hashMap {
		if !hashMap[num-1] {
			currentNum := num
			currentLen := 1
			for hashMap[currentNum+1] {
				currentNum++
				currentLen++
			}
			if currentLen > max {
				max = currentLen
			}
		}
	}
	return max
}
