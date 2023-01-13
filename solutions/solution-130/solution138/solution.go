package solution138

func copyRandomList(head *Node) *Node {
	ptr := head
	if ptr == nil {
		return nil
	}

	// 收集链表信息
	nodeList := make([]*Node, 0)
	nodeIndexMap := make(map[*Node]int)
	index := 1
	for ptr != nil {
		node := &Node{Val: ptr.Val}
		nodeList = append(nodeList, node)
		nodeIndexMap[ptr] = index

		ptr = ptr.Next
		index++
	}

	randList := make([]int, len(nodeList))
	ptr = head
	for i := 0; ptr != nil; i++ {
		if index := nodeIndexMap[ptr.Random]; index != 0 {
			randList[i] = index - 1
		} else {
			randList[i] = -1
		}
		ptr = ptr.Next
	}

	// 构建链表关系
	for i := range nodeList {
		if i != 0 {
			nodeList[i-1].Next = nodeList[i]
		}
		if randList[i] != -1 {
			nodeList[i].Random = nodeList[randList[i]]
		}
	}
	return nodeList[0]
}
