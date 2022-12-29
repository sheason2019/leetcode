package solution110

func connect2(root *Node) *Node {
	if root == nil {
		return root
	}
	layerScanner([]*Node{root})

	return root
}

func layerScanner2(layer []*Node) {
	nextLayer := []*Node{}

	for index, node := range layer {
		if node.Left != nil {
			nextLayer = append(nextLayer, node.Left)
		}
		if node.Right != nil {
			nextLayer = append(nextLayer, node.Right)
		}
		if index > 0 {
			layer[index-1].Next = node
		}
	}

	if len(nextLayer) > 0 {
		layerScanner(nextLayer)
	}
}
