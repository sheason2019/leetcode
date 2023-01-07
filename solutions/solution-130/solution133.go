package solution130

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	graphNodes := make(map[int]*Node)
	fetchGraphNodes(node, graphNodes)

	copyNodes := make(map[int]*Node)
	for v := range graphNodes {
		copyNodes[v] = &Node{
			Val:       v,
			Neighbors: []*Node{},
		}
	}
	for v, node := range graphNodes {
		for _, neightbor := range node.Neighbors {
			copyNodes[v].Neighbors = append(copyNodes[v].Neighbors, copyNodes[neightbor.Val])
		}
	}

	return copyNodes[1]
}

func fetchGraphNodes(node *Node, nodeMap map[int]*Node) {
	nodeMap[node.Val] = node
	for _, neighbor := range node.Neighbors {
		if nodeMap[neighbor.Val] == nil {
			fetchGraphNodes(neighbor, nodeMap)
		}
	}
}
