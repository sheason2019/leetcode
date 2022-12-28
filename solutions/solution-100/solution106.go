package solution100

func buildTree2(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 1 {
		return &TreeNode{Val: postorder[len(postorder)-1]}
	}

	node := TreeNode{Val: postorder[len(postorder)-1]}
	rootIndex := 0

	for i, v := range inorder {
		if v == node.Val {
			rootIndex = i
		}
	}

	inorder_left := inorder[:rootIndex]
	inorder_right := inorder[rootIndex+1:]

	postorder_left := postorder[:len(inorder_left)]
	postorder_right := postorder[len(inorder_left) : len(postorder)-1]

	if len(inorder_left) > 0 {
		node.Left = buildTree2(inorder_left, postorder_left)
	}
	if len(inorder_right) > 0 {
		node.Right = buildTree2(inorder_right, postorder_right)
	}

	return &node
}
