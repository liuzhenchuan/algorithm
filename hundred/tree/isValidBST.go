package tree

import "math"

// isValidBST 验证二叉搜索树 https://leetcode.cn/problems/validate-binary-search-tree/description/?envType=study-plan-v2&envId=top-100-liked
func isValidBST(root *TreeNode) bool {
	return helper1(root, math.MinInt64, math.MaxInt64)
}

func helper1(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper1(root.Left, lower, root.Val) && helper1(root.Right, root.Val, upper)
}
