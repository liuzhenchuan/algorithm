package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// inorderTraversal 二叉树中序遍历 https://leetcode.cn/problems/binary-tree-inorder-traversal/description/?envType=study-plan-v2&envId=top-100-liked
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return res
}

func inorderTraversal1(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, inorderTraversal1(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal1(root.Right)...)
	return res

}

type nodeWithColor struct {
	Color int
	Node  *TreeNode
}

func inorderTraversal2(root *TreeNode) []int {
	WHITE, GRAY := 0, 1
	var res = []int{}
	stack := make([]nodeWithColor, 0)
	stack = append(stack, nodeWithColor{
		Color: WHITE,
		Node:  root,
	})
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top.Node == nil {
			continue
		}
		if top.Color == WHITE {
			stack = append(stack, nodeWithColor{
				Color: WHITE,
				Node:  top.Node.Right,
			})
			stack = append(stack, nodeWithColor{
				Color: GRAY,
				Node:  top.Node,
			})
			stack = append(stack, nodeWithColor{
				Color: WHITE,
				Node:  top.Node.Left,
			})
		} else {
			res = append(res, top.Node.Val)
		}
	}
	return res
}
