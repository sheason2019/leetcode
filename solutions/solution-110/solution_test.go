package solution110

import (
	"fmt"
	"testing"

	"github.com/sheason2019/leetcode/utils"
)

const null = -999

func TestSolution110(t *testing.T) {
	tree := generateTree([]int{3, 9, 20, null, null, 15, 7})
	fmt.Println(isBalanced(tree))
}
func TestSolution110_1(t *testing.T) {
	tree := generateTree([]int{1, 2, 2, 3, 3, null, null, 4, 4})
	fmt.Println(isBalanced(tree))
}
func TestSolution110_2(t *testing.T) {
	tree := generateTree([]int{1, 2, 3, 4, 5, 6, null, 8})
	fmt.Println(isBalanced(tree))
}
func TestMinDepth_1(t *testing.T) {
	root := generateTree([]int{3, 9, 20, null, null, 15, 7})
	fmt.Println(minDepth(root))
}
func TestMinDepth_2(t *testing.T) {
	root := generateTree([]int{2, null, 3, null, 4, null, 5, null, 6})
	fmt.Println(minDepth(root))
}

func TestHasPathSum(t *testing.T) {
	root := generateTree([]int{5, 4, 8, 11, null, 13, 4, 7, 2, null, null, null, 1})
	fmt.Println(hasPathSum(root, 22))
}
func TestHasPathSum_2(t *testing.T) {
	root := generateTree([]int{-2, null, -3})
	fmt.Println(hasPathSum(root, -5))
}
func TestHasPathSum_3(t *testing.T) {
	root := generateTree([]int{1, -2, -3, 1, 3, -2, null, -1})
	fmt.Println(hasPathSum(root, -1))
}
func TestPathSum_1(t *testing.T) {
	root := generateTree([]int{5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1})
	fmt.Println(pathSum(root, 22))
}
func TestFlatten(t *testing.T) {
	root := generateTree([]int{1, 2, 5, 3, 4, null, 6})
	flatten(root)
	utils.JsonLog(root)
}
func TestNumDistinct(t *testing.T) {
	s := "rabbbit"
	ts := "rabbit"
	fmt.Println(numDistinct(s, ts))
}
func TestConnect(t *testing.T) {
	node := generateTreeNode([]int{1, 2, 3, 4, 5, 6, 7})
	ans := connect(node)
	utils.JsonLog(ans)
}
func TestGenerate(t *testing.T) {
	fmt.Println(generate(10))
}
func TestGetRow(t *testing.T) {
	fmt.Println(getRow(3))
}
