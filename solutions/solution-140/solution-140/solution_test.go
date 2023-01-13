package solution140

import (
	"fmt"
	"testing"
)

func TestWordBreak1(t *testing.T) {
	ans := wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"})
	for _, v := range ans {
		fmt.Println(v)
	}
}

func TestWordBreak2(t *testing.T) {
	ans := wordBreak("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"})
	for _, v := range ans {
		fmt.Println(v)
	}
}

func TestWordBreak3(t *testing.T) {
	ans := wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})
	for _, v := range ans {
		fmt.Println(v)
	}
}
