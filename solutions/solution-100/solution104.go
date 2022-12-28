package solution100

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 使用递归将每层的值写入从参数传入的二维数组即可

	return depthScanner([]*TreeNode{root}, 1)
}

func depthScanner(layer []*TreeNode, depth int) int {
	nextLayer := []*TreeNode{}
	for _, node := range layer {
		if node.Left != nil {
			nextLayer = append(nextLayer, node.Left)
		}
		if node.Right != nil {
			nextLayer = append(nextLayer, node.Right)
		}
	}

	if len(nextLayer) > 0 {
		return depthScanner(nextLayer, depth+1)
	}

	return depth
}
