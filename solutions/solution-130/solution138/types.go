package solution138

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

const null = -99999

func GenNode(source [][]int) *Node {
	nodeList := make([]*Node, len(source))
	for i, list := range source {
		nodeList[i] = &Node{
			Val: list[0],
		}
	}
	for i, list := range source {
		if list[1] != null {
			nodeList[i].Random = nodeList[list[1]]
		}
		if i != 0 {
			nodeList[i-1].Next = nodeList[i]
		}
	}

	return nodeList[0]
}

func PrintNode(node *Node) {
	indexMap := make(map[*Node]int)
	ptr := node
	for i := 1; ptr != nil; i++ {
		indexMap[ptr] = i
		ptr = ptr.Next
	}

	ptr = node
	fmt.Print("[")
	for ptr != nil {
		index := indexMap[ptr.Random]
		var random string
		if index == 0 {
			random = "null"
		} else {
			random = fmt.Sprint(index - 1)
		}

		fmt.Printf("[%v, %v]", ptr.Val, random)
		ptr = ptr.Next
		if ptr != nil {
			fmt.Printf(",")
		}
	}
	fmt.Print("]\n")
}
