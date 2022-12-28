package solution100

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	// 使用递归将每层的值写入从参数传入的二维数组即可
	values := [][]int{}
	levelScanner([]*TreeNode{root}, &values)

	return values
}

func levelScanner(layer []*TreeNode, values *[][]int) {
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
	*values = append(*values, layerValue)
	if len(nextLayer) > 0 {
		levelScanner(nextLayer, values)
	}
}
