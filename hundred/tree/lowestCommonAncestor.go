package tree

// lowestCommonAncestor 二叉树的最近公共祖先 https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/description/?envType=study-plan-v2&envId=top-100-liked
// 如果x节点是最近公共祖先，那么有如下两种情况
// 1、x左节点包含p或q， 且x右节点包含p或q
// 2、p和q同时在左边或同时在右边
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val { //找到p或q就返回该节点
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil { // 左边找到1个，右边找到1个，则root是最近公共祖先
		return root
	}
	if left == nil { // 仅一边找到，则该节点就是最近公共祖先
		return right
	}
	return left
}
