package link

/**
https://leetcode-cn.com/problems/reverse-nodes-in-k-group/
题目：给你一个链表，每k个节点一组进行翻转，请你返回翻转后的链表。
k是一个正整数，它的值小于或等于链表的长度。
如果节点总数不是k的整数倍，那么请将最后剩余的节点保持原有顺序。
进阶：
你可以设计一个只使用常数额外空间的算法来解决此问题吗？
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
解题思路：
1、定义hair节点，指向头节点，以及首次翻转后的头节点
2，pre永远指向待翻转链表head，翻转后再指向新head
3、只要head不为nil，遍历链表，先移动k个，如果小于k，直接返回hair节点，否则找到tail节点
3、翻转head到tail链表
4、将原来到pre接到翻转后到head，翻转后到tail接到原来待翻转的头
5、重复，直到不足k个
*/

func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head} // 定义一个hair节点，指向head
	pre := hair                   // 定义pre节点，初始指为hair，指向待翻转k段的head

	for head != nil { // 遍历链表
		tail := pre              // 定义尾节点，指向pre
		for i := 0; i < k; i++ { // 遍历前k个节点，找到第k个节点，赋值给tail
			tail = tail.Next
			if tail == nil { // 如果不到k个节点，直接返回hair.Next
				return hair.Next
			}
		}

		next := tail.Next                  // next标示下一段k个节点的第一个节点
		head, tail = myReverse(head, tail) // 反转一段k个节点，返回新的头节点和尾节点
		pre.Next = head                    // pre 接到翻转后到head （注意首次翻转时pre=hair,hair指向最终结果头节点）
		tail.Next = next                   // 翻转后到tail指向下一个待翻转的节点
		pre = tail                         // pre移到待翻转前一个节点 （第一次翻转后，pre和hair节点再此分离，hair永远指向头）
		head = tail.Next                   // head是待翻转节点头节点
	}

	return hair.Next // 翻转后头节点
}

func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next // 暂存k段之后段头节点
	p := head         // p节点初始head
	for prev != nil {
		nex := p.Next // 指向下一个节点
		p.Next = prev
		prev = p
		p = nex
	}
	return tail, head
}
