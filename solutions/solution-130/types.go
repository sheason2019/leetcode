package solution130

type Node struct {
	Val       int
	Neighbors []*Node
}

func GenGraphNode(edges [][]int) *Node {
	graphNodes := make(map[int]*Node)
	for val := range edges {
		node := &Node{
			Val:       val + 1,
			Neighbors: []*Node{},
		}
		graphNodes[val+1] = node
	}

	for val, neighbors := range edges {
		for _, neighbor := range neighbors {
			graphNodes[val+1].Neighbors = append(graphNodes[val+1].Neighbors, graphNodes[neighbor])
		}
	}

	return graphNodes[1]
}
