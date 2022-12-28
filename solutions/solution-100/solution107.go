package solution100

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	// 使用递归将每层的值写入从参数传入的二维数组即可
	values := [][]int{}
	levelScannerBottom([]*TreeNode{root}, &values)

	return values
}

func levelScannerBottom(layer []*TreeNode, values *[][]int) {
	layerValue := make([]int, len(layer))
	nextLayer := []*TreeNode{}
	for i, node := range layer {
		layerValue[i] = node.Val
		if node.Left != nil {
			nextLayer = append(nextLayer, node.Left)
		}
		if node.Right != nil {
			nextLayer = append(nextLayer, node.Right)
		}
	}
	if len(nextLayer) > 0 {
		levelScannerBottom(nextLayer, values)
	}
	*values = append(*values, layerValue)
}
