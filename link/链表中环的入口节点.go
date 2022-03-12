package link

/**
给定一个链表，返回链表开始入环的第一个节点。 从链表的头节点开始沿着 next 指针进入环的第一个节点为环的入口节点。如果链表无环，则返回null。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意，pos 仅仅是用于标识环的情况，并不会作为参数传递到函数中。

说明：不允许修改给定的链表。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/c32eOV
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
解法：
双指针法，fast每次走两步，slow每次走一步，当fast和slow相遇时，说明链表有环
设从head到环入口处为a,入口处到相遇点长度为b，相遇点到入口点长度为c
f=a+n(b+c)+b=a+(n+1)b+nc=2(a+b)=2a+2b
a=(n+1)b+nc-2b=(n-1)b+nc=(n-1)(b+c)+c 也就是说a等于从相遇点，再走n-1圈，然后再走c
那么从相遇时，新指针ptr从head开始走，slow指针从相遇点开始走，两个指针将会在环入口相遇，则得到结果
*/
func detectCycle(head *ListNode) *ListNode {
	// 判断为空
	if head == nil {
		return nil
	}
	// 定义fast、slow初始化为head
	fast, slow := head, head
	// 开始遍历，结束条件是fast的下一个为nil或者下下个为nil，这种表示无环，返回nil
	for fast.Next != nil && fast.Next.Next != nil {
		// fast走2步，slow走1步
		fast = fast.Next.Next
		slow = slow.Next
		// 如果fast==slow，相遇,有环
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}
