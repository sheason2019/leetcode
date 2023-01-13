package solution139

import (
	"fmt"
	"testing"
)

func TestWordBreak1(t *testing.T) {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
}

func TestWordBreak2(t *testing.T) {
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
}

func TestWordBreak3(t *testing.T) {
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}

func TestWordBreak4(t *testing.T) {
	fmt.Println(wordBreak("aaaaaaa", []string{"aa", "aaaa"}))
}
