package solution130

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	board1 := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'O', 'X', 'O', 'X'}}
	solve(board1)
	for _, row := range board1 {
		fmt.Printf("%c\n", row)
	}

	fmt.Println()

	board2 := [][]byte{{'O', 'X', 'X', 'O', 'X'}, {'X', 'O', 'O', 'X', 'O'}, {'X', 'O', 'X', 'O', 'X'}, {'O', 'X', 'O', 'O', 'O'}, {'X', 'X', 'O', 'X', 'O'}}
	solve(board2)
	for _, row := range board2 {
		fmt.Printf("%c\n", row)
	}

}

func TestPartition(t *testing.T) {
	s := "aab"
	fmt.Println(partition(s))
	fmt.Println(partition("a"))
}

func TestMinCut(t *testing.T) {
	s := "aab"
	fmt.Println(minCut(s))
	fmt.Println(minCut("apjesgpsxoeiokmqmfgvjslcjukbqxpsobyhjpbgdfruqdkeiszrlmtwgfxyfostpqczidfljwfbbrflkgdvtytbgqalguewnhvvmcgxboycffopmtmhtfizxkmeftcucxpobxmelmjtuzigsxnncxpaibgpuijwhankxbplpyejxmrrjgeoevqozwdtgospohznkoyzocjlracchjqnggbfeebmuvbicbvmpuleywrpzwsihivnrwtxcukwplgtobhgxukwrdlszfaiqxwjvrgxnsveedxseeyeykarqnjrtlaliyudpacctzizcftjlunlgnfwcqqxcqikocqffsjyurzwysfjmswvhbrmshjuzsgpwyubtfbnwajuvrfhlccvfwhxfqthkcwhatktymgxostjlztwdxritygbrbibdgkezvzajizxasjnrcjwzdfvdnwwqeyumkamhzoqhnqjfzwzbixclcxqrtniznemxeahfozp"))
}

func TestCloneGraph(t *testing.T) {
	graph := GenGraphNode([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}})
	newGraph := cloneGraph(graph)
	fmt.Printf("%+v\n", newGraph)
}

func TestCanCompleteCircuit(t *testing.T) {
	canCompleteCircuit([]int{}, []int{})
}

func TestCandy(t *testing.T) {
	fmt.Println(candy([]int{1, 0, 2}))
	fmt.Println(candy([]int{1, 2, 2}))
}

func TestSingleNumber(t *testing.T) {
	fmt.Println(singleNumber([]int{2, 2, 1}))
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
}

func TestSingleNumber2(t *testing.T) {
	fmt.Println(singleNumber2([]int{1, 1, 1, 3}))
}
