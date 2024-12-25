package link

// removeNthFromEnd 删除链表倒数第N个元素
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l := getLength(head)
	dummy := &ListNode{0, head}
	curr := dummy
	for i := 0; i < l-n; i++ {
		curr = curr.Next
	}
	curr.Next = curr.Next.Next
	return dummy.Next
}

func getLength(head *ListNode) int {
	l := 0
	for head != nil {
		l++
		head = head.Next
	}
	return l
}

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	nodes := []*ListNode{}
	dummy := &ListNode{0, head}
	for node := dummy; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	prev := nodes[len(nodes)-1-n]
	prev.Next = prev.Next.Next
	return dummy.Next
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := head, dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for ; first != nil; first = first.Next {
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}
