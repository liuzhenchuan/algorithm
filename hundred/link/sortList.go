package link

// sortList 排序链表 https://leetcode.cn/problems/sort-list/description/?envType=study-plan-v2&envId=top-100-liked
func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}
	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	mid := slow
	return merge(sort(head, mid), sort(mid, tail))
}

func merge(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp, temp1, temp2 := dummyHead, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val < temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	}
	if temp2 != nil {
		temp.Next = temp2
	}
	return dummyHead.Next
}
