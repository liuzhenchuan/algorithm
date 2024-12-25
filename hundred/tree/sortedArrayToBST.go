package tree

// sortedArrayToBST 将有序数组转换为二叉搜索树 https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/description/?envType=study-plan-v2&envId=top-100-liked
func sortedArrayToBST(nums []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	root := &TreeNode{
		Val: nums[mid],
	}
	root.Left = helper(nums, left, mid-1)
	root.Right = helper(nums, mid+1, right)
	return root
}
