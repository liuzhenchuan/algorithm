package tree

/*
https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/
题目：给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）
解题思路：
基本和二叉树层序遍历相同。只需在奇数层时，将数据翻转
*/

func ZigzagLevelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}                // 将第一层的root放到队列
	for level := 0; len(queue) > 0; level++ { // level 标示每层， 队列不能为空
		vals := []int{} // 临时切片，存储节点值
		q := queue      // 当前层队列
		queue = nil
		for _, node := range q { // 出当前层队列
			vals = append(vals, node.Val) // 出队列的值放到vals中
			if node.Left != nil {
				queue = append(queue, node.Left) // 将左节点放到队列
			}
			if node.Right != nil {
				queue = append(queue, node.Right) // 将右节点放入队列
			}
		}
		// 本质上和层序遍历一样，我们只需要把奇数层的元素翻转即可
		if level%2 == 1 {
			for i, n := 0, len(vals); i < n/2; i++ {
				vals[i], vals[n-1-i] = vals[n-1-i], vals[i]
			}
		}
		ans = append(ans, vals)
	}
	return ans
}
