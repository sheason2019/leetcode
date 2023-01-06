package solution120

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const null = -999

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
