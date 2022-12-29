package solution110

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func generateTree(values []int) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	head := &TreeNode{Val: values[0]}
	values = values[1:]

	nodeBuf := []*TreeNode{head}
	for len(values) > 0 {
		nextNodeBuf := []*TreeNode{}
		for _, node := range nodeBuf {
			if values[0] != -999 {
				node.Left = &TreeNode{Val: values[0]}
				nextNodeBuf = append(nextNodeBuf, node.Left)
			}
			values = values[1:]
			if len(values) == 0 {
				break
			}
			if values[0] != -999 {
				node.Right = &TreeNode{Val: values[0]}
				nextNodeBuf = append(nextNodeBuf, node.Right)
			}
			values = values[1:]
			if len(values) == 0 {
				break
			}
		}
		nodeBuf = nextNodeBuf
	}

	return head
}

func generateTreeNode(values []int) *Node {
	if len(values) == 0 {
		return nil
	}

	head := &Node{Val: values[0]}
	values = values[1:]

	nodeBuf := []*Node{head}
	for len(values) > 0 {
		nextNodeBuf := []*Node{}
		for _, node := range nodeBuf {
			if values[0] != -999 {
				node.Left = &Node{Val: values[0]}
				nextNodeBuf = append(nextNodeBuf, node.Left)
			}
			values = values[1:]
			if len(values) == 0 {
				break
			}
			if values[0] != -999 {
				node.Right = &Node{Val: values[0]}
				nextNodeBuf = append(nextNodeBuf, node.Right)
			}
			values = values[1:]
			if len(values) == 0 {
				break
			}
		}
		nodeBuf = nextNodeBuf
	}

	return head
}
