package solution120

import (
	"fmt"
	"testing"
)

func TestMinimumTotal(t *testing.T) {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println(minimumTotal(triangle))
}

func TestMaxProfit(t *testing.T) {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}

func TestMaxProfit1(t *testing.T) {
	fmt.Println(maxProfit1([]int{7, 1, 5, 3, 6, 4}))
}

func TestMaxProfit2(t *testing.T) {
	fmt.Println(maxProfit2([]int{3, 3, 5, 0, 0, 3, 1, 4}))
}

func TestMaxPathSum(t *testing.T) {
	fmt.Println(maxPathSum(generateTree([]int{1, 2, 3})))
	fmt.Println(maxPathSum(generateTree([]int{-10, 9, 20, null, null, 15, 7})))
}

func TestIsPalindrom(t *testing.T) {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}
