package tree

// diameterOfBinaryTree 二叉树的直径 https://leetcode.cn/problems/diameter-of-binary-tree/description/?envType=study-plan-v2&envId=top-100-liked
func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	var depth func(node *TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftDepth := depth(node.Left)
		rightDepth := depth(node.Right)
		if leftDepth+rightDepth > maxDiameter {
			maxDiameter = leftDepth + rightDepth
		}
		return max(leftDepth, rightDepth) + 1
	}
	depth(root)
	return maxDiameter
}
