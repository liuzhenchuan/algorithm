package tree

// pathSum 路径总合 https://leetcode.cn/problems/path-sum-iii/description/?envType=study-plan-v2&envId=top-100-liked
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := rootSum(root, targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res
}

func rootSum(root *TreeNode, targetSum int) (res int) {
	if root == nil {
		return 0
	}
	val := root.Val
	if val == targetSum {
		res++
	}
	res += rootSum(root.Left, targetSum-val)
	res += rootSum(root.Right, targetSum-val)
	return
}
