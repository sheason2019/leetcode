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
		for runes[left] == ' ' {
			left++
		}
		for runes[right] == ' ' {
			right--
		}
		if left < right {
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
