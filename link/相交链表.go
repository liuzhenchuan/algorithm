package link

/**
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null
https://leetcode-cn.com/problems/intersection-of-two-linked-lists/

解题思路：
哈希集合
遍历A，将节点都存在map中，遍历B，没遍历一个节点检查是否在map中，是则返回该节点
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 如果任意一个链表为空，返回nil
	if headA == nil || headB == nil {
		return nil
	}
	// 定义map
	m := make(map[*ListNode]bool)
	// 遍历A，将所有节点都存到map中
	p := headA
	for p != nil {
		m[p] = true
		p = p.Next
	}
	// 遍历B，每次检查是否在map中，是返回该节点
	p = headB
	for p != nil {
		_, ok := m[p]
		if ok {
			return p
		}
		p = p.Next
	}
	return nil
}

/*
解题思路：
双指针
指针同时向后遍历，当A到尾后，指针移到B的头，当B到尾后，指针移到A的头，直到相交
证明相交的情况：
假设A不相交部分长度为a，B不相交部分长度为b，相交部分为c
A走完的路程是a+c,再从B的头走到相交节点，总共路程是a+c+b
B走完的路程是b+c,再从A的头走到相交点，总共路程是b+c+a  AB路程相等
证明不相交的情况：
假设A长度为a，B长度为b
A走完的路程是a，再从B的头走完，总共是a+b
B走完的路程是b，再从A的头走完，总共是b+a  AB同时走完
*/
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	// 如果任意一个链表为空，返回nil
	if headA == nil || headB == nil {
		return nil
	}
	a := headA
	b := headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}
