package solution110

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	layerScanner([]*Node{root})

	return root
}

func layerScanner(layer []*Node) {
	nextLayer := []*Node{}

	hasChild := layer[0].Left != nil
	for index, node := range layer {
		if hasChild {
			nextLayer = append(nextLayer, node.Left)
			nextLayer = append(nextLayer, node.Right)
		}
		if index > 0 {
			layer[index-1].Next = node
		}
	}

	if hasChild {
		layerScanner(nextLayer)
	}
}
