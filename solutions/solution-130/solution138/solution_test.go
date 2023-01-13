package solution138

import (
	"testing"
)

func TestCopyRandomList1(t *testing.T) {
	head := GenNode([][]int{{7, null}, {13, 0}, {11, 4}, {10, 2}, {1, 0}})
	cpHead := copyRandomList(head)
	PrintNode(head)
	PrintNode(cpHead)
}

func TestCopyRandomList2(t *testing.T) {
	head := GenNode([][]int{{7, null}, {13, 0}, {11, 4}, {10, 2}, {1, 0}})
	cpHead := copyRandomList(head)
	PrintNode(head)
	PrintNode(cpHead)
}

func TestCopyRandomList3(t *testing.T) {
	head := GenNode([][]int{{7, null}, {13, 0}, {11, 4}, {10, 2}, {1, 0}})
	cpHead := copyRandomList(head)
	PrintNode(head)
	PrintNode(cpHead)
}
