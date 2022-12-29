package solution100

import (
	"fmt"
	"testing"

	"github.com/sheason2019/leetcode/utils"
)

func createTree() *TreeNode {
	head := TreeNode{
		Val: 3,
	}
	head.Left = &TreeNode{Val: 9}
	head.Right = &TreeNode{
		Val:   20,
		Left:  &TreeNode{Val: 15},
		Right: &TreeNode{Val: 7},
	}

	return &head
}

func TestLevelOrder(t *testing.T) {
	head := createTree()
	values := levelOrder(head)
	fmt.Println(values)
}

func TestZigzagLevelOrder(t *testing.T) {
	head := createTree()
	values := zigzagLevelOrder(head)
	fmt.Println(values)
}

func TestMaxDepth(t *testing.T) {
	head := createTree()
	depth := maxDepth(head)
	fmt.Println(depth)
}

func TestSlice(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(slice[6:])
}

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	node := buildTree(preorder, inorder)
	utils.JsonLog(node)
}
func TestBuildTree2(t *testing.T) {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}

	node := buildTree2(inorder, postorder)
	utils.JsonLog(node)
}
func TestLevelOrderBottom(t *testing.T) {
	head := createTree()
	values := levelOrderBottom(head)
	fmt.Println(values)
}

func TestSortedArrayToBST(t *testing.T) {
	nums := []int{-10, -3, 0, 5, 9}
	node := sortedArrayToBST(nums)

	utils.JsonLog(node)
}

func createList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := ListNode{Val: values[0]}
	if len(values) == 1 {
		return &head
	}

	ptr := &head
	for i := 1; i < len(values); i++ {
		ptr.Next = &ListNode{Val: values[i]}
		ptr = ptr.Next
	}

	return &head
}

func TestSortedListToBST(t *testing.T) {
	list := createList([]int{-10, -3, 0, 5, 9})
	node := sortedListToBST(list)

	utils.JsonLog(node)
}
