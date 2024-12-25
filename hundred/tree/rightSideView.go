package tree

// rightSideView 二叉树的右视图 https://leetcode.cn/problems/binary-tree-right-side-view/description/?envType=study-plan-v2&envId=top-100-liked
func rightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	stack := []*TreeNode{}
	stack = append(stack, root)
	for len(stack) > 0 {
		size := len(stack)
		for i := 0; i < size; i++ {
			node := stack[0]
			stack = stack[1:]
			if i == size-1 {
				res = append(res, node.Val)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
	}
	return res
}
