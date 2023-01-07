package solution120

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	graphHead := genGraph(beginWord, wordList)
	var graphEnd *GraphNode

	attachMap := make(map[*GraphNode]bool)
	attachMap[graphHead] = true

	// Graph分层
	layer := 0
	for len(attachMap) > 0 && graphEnd == nil {
		layer++

		nextMap := make(map[*GraphNode]bool)
		for node := range attachMap {
			if node.Layer == 0 {
				node.Layer = layer
			}
			if node.Val == endWord {
				graphEnd = node
				break
			}
			for _, edge := range node.Edges {
				if edge.Layer == 0 {
					nextMap[edge] = true
				}
			}
		}
		attachMap = nextMap
	}

	ansList := [][]string{}
	if graphEnd == nil {
		return ansList
	}

	ansNodes := [][]*GraphNode{{graphEnd}}
	for layer > 1 {
		layer--

		nextAnsNodes := [][]*GraphNode{}
		for _, option := range ansNodes {
			for _, edge := range option[len(option)-1].Edges {
				if edge.Layer == layer {
					nextNodes := []*GraphNode{}
					nextNodes = append(nextNodes, option...)
					nextNodes = append(nextNodes, edge)
					nextAnsNodes = append(nextAnsNodes, nextNodes)
				}
			}
		}
		ansNodes = nextAnsNodes
	}

	for i := range ansNodes {
		list := make([]string, len(ansNodes[i]))
		for j := range ansNodes[i] {
			list[j] = ansNodes[i][len(ansNodes[i])-1-j].Val
		}
		ansList = append(ansList, list)
	}

	return ansList
}

type GraphNode struct {
	Val   string
	Layer int
	Edges []*GraphNode
}

func genGraph(beginWord string, wordList []string) *GraphNode {
	var head *GraphNode

	addBeginWord := false
	// 生成graph node list
	nodes := []*GraphNode{}
	for _, word := range wordList {
		node := &GraphNode{
			Val:   word,
			Edges: []*GraphNode{},
		}
		if word == beginWord {
			addBeginWord = true
			head = node
		}
		nodes = append(nodes, node)
	}

	if !addBeginWord {
		head = &GraphNode{
			Val:   beginWord,
			Edges: []*GraphNode{},
		}
		nodes = append(nodes, head)
	}

	// 遍历list生成graph
	length := len(nodes)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if allowUse(nodes[i].Val, nodes[j].Val) {
				nodes[i].Edges = append(nodes[i].Edges, nodes[j])
				nodes[j].Edges = append(nodes[j].Edges, nodes[i])
			}
		}
	}

	return head
}

// 单词准入函数
func allowUse(source, target string) bool {
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

// 回溯算法 - 超时

// func findLadders(beginWord string, endWord string, wordList []string) [][]string {
// 	ansList := [][]string{}
// 	findLaddersResolve(beginWord, endWord, wordList, nil, &ansList)

// 	ansList = findShortestAns(&ansList)

// 	return ansList
// }

// // 回溯思路
// func findLaddersResolve(currentWord string, targetWord string, wordList []string, ans []string, ansList *[][]string) {
// 	if ans == nil {
// 		ans = []string{}
// 	}
// 	if currentWord == targetWord {
// 		ans = append(ans, currentWord)
// 		*ansList = append(*ansList, ans)
// 		return
// 	}

// 	for i, word := range wordList {
// 		if len(wordList) > 0 && allowUse(currentWord, word) {
// 			nextWordList := append(append([]string{}, wordList[:i]...), wordList[i+1:]...)
// 			nextAns := append(append([]string{}, ans...), currentWord)
// 			findLaddersResolve(word, targetWord, nextWordList, nextAns, ansList)
// 		}
// 	}
// }

// // 单词准入函数
// func allowUse(source, target string) bool {
// 	count := 0
// 	for i := 0; i < len(source); i++ {
// 		if source[i] != target[i] {
// 			count++
// 			if count > 1 {
// 				break
// 			}
// 		}
// 	}
// 	return count == 1
// }

// // 提取最短结果
// func findShortestAns(ansList *[][]string) [][]string {
// 	lengthMap := make(map[int][][]string)
// 	min := 600
// 	for _, ans := range *ansList {
// 		length := len(ans)
// 		if lengthMap[length] == nil {
// 			lengthMap[length] = [][]string{}
// 		}
// 		lengthMap[length] = append(lengthMap[length], ans)
// 		if length < min {
// 			min = length
// 		}
// 	}

// 	return lengthMap[min]
// }
