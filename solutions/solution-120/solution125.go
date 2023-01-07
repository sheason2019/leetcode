package solution120

import (
	"unicode"
)

func isPalindrome(s string) bool {
	runes := []rune(s)
	// 双指针
	left := 0
	right := len(s) - 1
	for left < right {
		for !isWord(runes[left]) {
			left++
			if left >= right {
				break
			}
		}
		for !isWord(runes[right]) {
			right--
			if left >= right {
				break
			}
		}
		if left >= right {
			break
		}

		if unicode.ToLower(runes[left]) != unicode.ToLower(runes[right]) {
			return false
		}

		left++
		right--
	}

	return true
}

func isWord(r rune) bool {
	if r >= 'a' && r <= 'z' {
		return true
	}
	if r >= 'A' && r <= 'Z' {
		return true
	}
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}
