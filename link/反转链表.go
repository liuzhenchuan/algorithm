package link

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
https://leetcode-cn.com/problems/reverse-linked-list/
题目：给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
解题思路：
1、current指针遍历链表
2、temp指针记录current节点，然后current指针向后移一位
3、将temp指针指向翻转头节点
4、翻转头节点向后移一位，等于temp，因为翻转后该节点就是翻转后的头节点
*/

func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pReverseHead *ListNode
	curr := head
	for curr != nil {
		temp := curr
		curr = curr.Next
		temp.Next = pReverseHead
		pReverseHead = temp
	}
	return pReverseHead
}

/*
 1、用curr指针遍历所有节点
 2、将当前节点指向前一个节点后，因为无法找到下一个节点，所以用next记录下一个节点
 3、当前节点指向前一个节点，因为不知道前一个节点是什么，所以用prev记录前一个节点
 4、把curr指针向后移一位，知道curr不为空
*/

func ReverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curr := head
	var prev *ListNode
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

/*
 1、用递归的思想处理
 2、把后n-1个节点先翻转，然后把下一个节点的next指向head
 3、把head指向空
*/

func Recursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev := Recursion(head.Next)
	head.Next.Next = head
	head.Next = nil
	return prev
}

/**
自创方法：
1、记录pre节点，cur节点，next节点
2、初始化pre是第一个节点，cur是第二节点
3、遍历cur节点，从第二个开始，直到不为空
4、记录next节点
5、把cur节点指向pre
6、pre向后移1位
7、cur向后移1位，但如果后面是nil，则返回cur
8、循环，返回cur
*/

func ReverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre := head
	pre.Next = nil
	cur := head.Next
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		if next == nil {
			return cur
		}
		cur = next
	}
	return cur
}
