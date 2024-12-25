package tree

func flatten(root *TreeNode) {
	list := preorderTraversal(root)
	for i := 1; i < len(list); i++ {
		prev := list[i-1]
		curr := list[i]
		prev.Left = nil
		prev.Right = curr
	}
}

func preorderTraversal(root *TreeNode) []*TreeNode {
	list := []*TreeNode{}
	if root != nil {
		list = append(list, root)
		list = append(list, preorderTraversal(root.Left)...)
		list = append(list, preorderTraversal(root.Right)...)
	}
	return list
}
