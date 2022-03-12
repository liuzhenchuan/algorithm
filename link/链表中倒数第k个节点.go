package link

/**
输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。

例如，一个链表有 6 个节点，从头节点开始，它们的值依次是 1、2、3、4、5、6。这个链表的倒数第 3 个节点是值为 4 的节点。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

双指针法：
1、初始化fast、slow两个指针
2、fast指针向前走k步，然后slow和fast一起往前走，直到fast到结尾，返回slow
3、注意处理，k大于链表长度的情况，直接返回nil
*/

func getKthFromEnd(head *ListNode, k int) *ListNode {
	// 判断链表是否为空，为空返回nil
	if head == nil {
		return nil
	}
	// 初始化fast、slow为head
	fast, slow := head, head
	// 定义fast走的步数i，从第1个开始
	i := 1
	// 开始遍历fast，直到fast为nil，
	for fast != nil {
		// 如果是最后一个节点，1 i<k 返回nil， i >= k, 返回slow
		if fast.Next == nil {
			if i < k {
				return nil
			} else {
				return slow
			}
		}
		// fast向后走一步，i加1
		fast = fast.Next
		i++
		// 判断i和k, 如果大于，slow向前走一步
		if i > k {
			slow = slow.Next
		}
	}
	return nil
}
