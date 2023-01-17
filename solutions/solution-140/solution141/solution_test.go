package solution141

import (
	"fmt"
	"testing"

	listnode "github.com/sheason2019/leetcode/types/list-node"
)

func TestHasCycle1(t *testing.T) {
	head := listnode.GenListNode([]int{3, 2, 0, -4}, 1)
	fmt.Println(hasCycle(head))
}

func TestHasCycle2(t *testing.T) {
	head := listnode.GenListNode([]int{1, 2}, 0)
	fmt.Println(hasCycle(head))
}

func TestHasCycle3(t *testing.T) {
	head := listnode.GenListNode([]int{1}, -1)
	fmt.Println(hasCycle(head))
}
