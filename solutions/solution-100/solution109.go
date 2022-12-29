package solution100

func sortedListToBST(head *ListNode) *TreeNode {
	// 链表长度为0时的情况
	if head == nil {
		return nil
	}

	// 快慢指针确定中点
	midNode := pointRace(head)
	right_head := midNode.Next

	node := TreeNode{
		Val: midNode.Val,
	}

	if right_head != nil {
		node.Right = sortedListToBST(right_head)
	}
	if midNode != head {
		node.Left = sortedListToBST(head)
	}

	// 然后用两个新指针拿到左右子树的Node
	return &node
}

// 快慢指针竞速，传入头节点，返回中间节点
func pointRace(head *ListNode) *ListNode {
	// last用来指slowPtr前面的节点，用来切断链表
	var last *ListNode
	slowPtr := head
	fastPtr := head

	for {
		for i := 0; i < 2; i++ {
			fastPtr = fastPtr.Next
			if fastPtr == nil {
				if last != nil {
					last.Next = nil
				}
				return slowPtr
			}
		}
		last = slowPtr
		slowPtr = slowPtr.Next
	}
}
