package tree

// invertTree 翻转二叉树 https://leetcode.cn/problems/invert-binary-tree/description/?envType=study-plan-v2&envId=top-100-liked
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}
