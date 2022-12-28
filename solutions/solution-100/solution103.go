package solution100

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	values := [][]int{}
	zigzagLayerScanner([]*TreeNode{root}, &values, true)

	return values
}

func zigzagLayerScanner(layer []*TreeNode, values *[][]int, isLtr bool) {
	nextLayer := []*TreeNode{}
	layerValue := make([]int, len(layer))

	for i, node := range layer {
		layerValue[i] = node.Val
		if node.Left != nil {
			nextLayer = append(nextLayer, node.Left)
		}
		if node.Right != nil {
			nextLayer = append(nextLayer, node.Right)
		}
	}

	if !isLtr {
		// reverse value
		for i, j := 0, len(layerValue)-1; i < j; i, j = i+1, j-1 {
			layerValue[i], layerValue[j] = layerValue[j], layerValue[i]
		}
	}

	*values = append(*values, layerValue)
	if len(nextLayer) > 0 {
		zigzagLayerScanner(nextLayer, values, !isLtr)
	}
}
