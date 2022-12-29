package solution110

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	var left_depth int = 999999
	var right_depth int = 999999

	if root.Left != nil {
		left_depth = minDepth(root.Left)
	}
	if root.Right != nil {
		right_depth = minDepth(root.Right)
	}

	if left_depth < right_depth {
		return left_depth + 1
	} else {
		return right_depth + 1
	}
}
