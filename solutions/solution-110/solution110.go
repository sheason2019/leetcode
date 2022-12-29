package solution110

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return treeDepth(root) != -1
}

func treeDepth(root *TreeNode) int {
	var left_depth int
	var right_depth int

	if root.Left != nil {
		left_depth = treeDepth(root.Left)
	}
	if root.Right != nil {
		right_depth = treeDepth(root.Right)
	}
	if left_depth == -1 || right_depth == -1 {
		return -1
	}

	// 如果左右子树高度差大于1，返回-1表示不平衡
	if (left_depth-right_depth)*(left_depth-right_depth) > 1 {
		return -1
	}

	// 否则返回该节点的高度
	if left_depth > right_depth {
		return left_depth + 1
	} else {
		return right_depth + 1
	}
}
