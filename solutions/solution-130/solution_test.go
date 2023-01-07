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
