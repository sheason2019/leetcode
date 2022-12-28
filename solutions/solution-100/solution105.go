package solution100

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}

	node := TreeNode{Val: preorder[0]}
	rootIndex := 0

	for i, v := range inorder {
		if v == node.Val {
			rootIndex = i
		}
	}

	inorder_left := inorder[:rootIndex]
	inorder_right := inorder[rootIndex+1:]

	preorder_left := preorder[1 : 1+len(inorder_left)]
	preorder_right := preorder[1+len(inorder_left):]

	if len(inorder_left) > 0 {
		node.Left = buildTree(preorder_left, inorder_left)
	}
	if len(inorder_right) > 0 {
		node.Right = buildTree(preorder_right, inorder_right)
	}

	return &node
}
