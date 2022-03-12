package tree

/*
https://leetcode-cn.com/problems/binary-tree-level-order-traversal/
题目：给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）
解题思路：
1、将每层的节点放到队列，第一层root
2、出队列，判断左节点是否为空，如果不为空，把左节点加入下一层队列，如果右节点不为空，把右节点加入队列
3、当前队列出完之后，开始遍历下一层队列，直到最后一层
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{} //初始化返回值，第一个层，第二个为每层节点值
	if root == nil { //如果为空，直接返回
		return ret
	}
	q := []*TreeNode{root}        //定义一个存储节点的队列，并把root节点放到队列
	for i := 0; len(q) > 0; i++ { //遍历没层，i表示层，必须队列不能为空
		ret = append(ret, []int{})    //结果中加入一个空值
		p := []*TreeNode{}            // 定义一个切片，存储节点
		for j := 0; j < len(q); j++ { // j 标示队列的头，遍历队列
			node := q[j]                      // 队列第一个节点
			ret[i] = append(ret[i], node.Val) // 将队列第j个值，放到第i层后
			if node.Left != nil {             // 如果取出的节点左节点不为空
				p = append(p, node.Left) // 把左节点放到p队列
			}
			if node.Right != nil {
				p = append(p, node.Right) // 把右节点放到队列
			}
		}
		q = p // 上层队列出完之后，把下层队列赋值给q，再次遍历
	}
	return ret
}
