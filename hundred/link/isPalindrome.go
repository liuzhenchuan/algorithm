package link

// isPalindrome 回文链表 https://leetcode.cn/problems/palindrome-linked-list/description/?envType=study-plan-v2&envId=top-100-liked
func isPalindrome(head *ListNode) bool {
	vals := []int{}
	for ; head != nil; head = head.Next {
		vals = append(vals, head.Val)
	}
	n := len(vals)
	for i, v := range vals[0 : n/2] {
		if v != vals[n-1-i] {
			return false
		}
	}
	return true
}
