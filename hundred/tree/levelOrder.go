package tree

// levelOrder 二叉树层序遍历 https://leetcode.cn/problems/binary-tree-level-order-traversal/description/?envType=study-plan-v2&envId=top-100-liked
func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}

	q := []*TreeNode{root}

	for len(q) > 0 {
		l := len(q)
		layer := []int{}
		for i := 0; i < l; i++ {
			top := q[0]
			q = q[1:]
			layer = append(layer, top.Val)
			if top.Left != nil {
				q = append(q, top.Left)
			}
			if top.Right != nil {
				q = append(q, top.Right)
			}
		}
		ret = append(ret, layer)
	}
	return ret
}
