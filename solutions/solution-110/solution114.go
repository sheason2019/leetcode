package solution110

func flatten(root *TreeNode) {
	flattenRecursion(root, &root)
}

func flattenRecursion(node *TreeNode, mountPtr **TreeNode) {
	if node == nil {
		return
	}

	left := node.Left
	right := node.Right

	node.Left = nil
	node.Right = nil

	if *mountPtr != node {
		(*mountPtr).Right = node
		*mountPtr = (*mountPtr).Right
	}

	flattenRecursion(left, mountPtr)
	flattenRecursion(right, mountPtr)
}
