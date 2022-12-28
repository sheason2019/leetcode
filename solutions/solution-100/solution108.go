package solution100

func sortedArrayToBST(nums []int) *TreeNode {
	// 使用递归创建树节点
	rootIndex := len(nums) / 2
	node := TreeNode{Val: nums[rootIndex]}

	nums_left := nums[:rootIndex]
	nums_right := nums[rootIndex+1:]

	if len(nums_left) > 0 {
		node.Left = sortedArrayToBST(nums_left)
	}
	if len(nums_right) > 0 {
		node.Right = sortedArrayToBST(nums_right)
	}

	return &node
}
