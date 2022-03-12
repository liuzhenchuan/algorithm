/**
  将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
  https://leetcode-cn.com/problems/merge-two-sorted-lists/
*/

package link

/**
递归法
1、判断两个节点大小，选出小的节点，然后小节点->剩下两个链表的merge结果即是最终结果
2、递归判出条件，如果都为空，返回空，任意一个为空，返回另外一个链表
*/

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 先判断出递归条件
	if list1 == nil && list2 == nil {
		return nil
	} else if list1 == nil && list2 != nil {
		return list2
	} else if list1 != nil && list2 == nil {
		return list1
	} else {
		if list1.Val <= list2.Val {
			list1.Next = mergeTwoLists(list1.Next, list2)
			return list1
		} else {
			list2.Next = mergeTwoLists(list1, list2.Next)
			return list2
		}
	}
}

/**
遍历法
1、找到比较小的头节点，作为返回节点，定义个tail节点，指向合并后的尾节点
2、cur1、cur2分别指向两个待合并链表进行遍历
3、任意找1条链表进行遍历，判断条件cur1!=nil && cur2!=nil,循环结束后，将tail指向不为空的cur
4、如果cur1<cur2，tail->cur1,tail=cur1,cur1向后一步
5、如果cur1=cur2, tail->cur1,tail=cur1,cur1向后一步，tail->cur2,tail=cur2,cur2向后一步
6、如果cur2>cur2, tail->cur2,tail=cur2,cur2向后一步
7、循环结束后将tail指向不为空的cur
*/
func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	} else if list1 == nil && list2 != nil {
		return list2
	} else if list1 != nil && list2 == nil {
		return list1
	} else {
		head := &ListNode{}
		if list1.Val <= list2.Val {
			head.Next = list1
		} else {
			head.Next = list2
		}
		tail := head
		cur1 := list1
		cur2 := list2
		for cur1 != nil && cur2 != nil {
			if cur1.Val < cur2.Val {
				tail.Next = cur1
				tail = cur1
				cur1 = cur1.Next
			} else if cur1.Val == cur2.Val {
				tail.Next = cur1
				tail = cur1
				cur1 = cur1.Next
				tail.Next = cur2
				tail = cur2
				cur2 = cur2.Next
			} else {
				tail.Next = cur2
				tail = cur2
				cur2 = cur2.Next
			}
		}
		if cur1 != nil {
			tail.Next = cur1
		}
		if cur2 != nil {
			tail.Next = cur2
		}
		return head.Next
	}
}
