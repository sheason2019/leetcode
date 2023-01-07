package solution120

func ladderLength(beginWord string, endWord string, wordList []string) int {
	graphHead, graphEnd := genNodes(beginWord, endWord, wordList)
	if graphEnd == nil {
		return 0
	}

	attachMap := make(map[*Node]bool)
	attachMap[graphHead] = true

	// Graph分层
	layer := 0
	for len(attachMap) > 0 {
		layer++

		nextMap := make(map[*Node]bool)
		for node := range attachMap {
			if node.Layer == 0 {
				node.Layer = layer
			}
			if node.Val == endWord {
				return layer
			}
			for _, edge := range node.Edges {
				if edge.Layer == 0 {
					nextMap[edge] = true
				}
			}
		}
		attachMap = nextMap
	}

	return 0
}

type Node struct {
	Val   string
	Layer int
	Edges []*Node
}

func genNodes(beginWord, endWord string, wordList []string) (head *Node, end *Node) {
	addBeginWord := false
	// 生成graph node list
	nodes := []*Node{}
	for _, word := range wordList {
		node := &Node{
			Val:   word,
			Edges: []*Node{},
		}
		if word == beginWord {
			addBeginWord = true
			head = node
		}
		if word == endWord {
			end = node
		}
		nodes = append(nodes, node)
	}

	if !addBeginWord {
		head = &Node{
			Val:   beginWord,
			Edges: []*Node{},
		}
		nodes = append(nodes, head)
	}

	// 遍历list生成graph
	length := len(nodes)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if allowEdge(nodes[i].Val, nodes[j].Val) {
				nodes[i].Edges = append(nodes[i].Edges, nodes[j])
				nodes[j].Edges = append(nodes[j].Edges, nodes[i])
			}
		}
	}

	return
}

// 单词准入函数
func allowEdge(source, target string) bool {
	count := 0
	for i := 0; i < len(source); i++ {
		if source[i] != target[i] {
			count++
			if count > 1 {
				break
			}
		}
	}
	return count == 1
}
