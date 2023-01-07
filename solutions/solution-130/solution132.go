package solution130

import "math"

func minCut(s string) int {
	// 字符串长度
	n := len(s)
	// DP[i][j]表示s[i:j]是否为回文字符串
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		for j := range dp[i] {
			dp[i][j] = true
		}
	}

	// 中心扩展法进行回文判断
	for l := n - 1; l >= 0; l-- {
		for r := l + 1; r < n; r++ {
			dp[l][r] = s[l] == s[r] && dp[l+1][r-1]
		}
	}

	// 创建另一个动态规划数组表示最少分割数量
	f := make([]int, n)
	for i := range f {
		// 若s[0:i]是回文，则无需分割
		if dp[0][i] {
			continue
		}
		f[i] = math.MaxInt
		for j := 0; j < i; j++ {
			if dp[j+1][i] && f[j]+1 < f[i] {
				f[i] = f[j] + 1
			}
		}
	}

	return f[n-1]
}

// 回溯失败 - 超时
// func minCut(s string) int {
// 	// 记录当前是第几次分割
// 	phase := 0
// 	strs := []string{s}
// 	for {
// 		nextStrs := []string{}
// 		for _, str := range strs {
// 			for length := len(str); length > 0; length-- {
// 				left := str[:length]
// 				right := str[length:]
// 				if checkPlalindrome(left) {
// 					if checkPlalindrome(right) {
// 						return phase
// 					}
// 					nextStrs = append(nextStrs, right)
// 				}
// 				if phase == 0 {
// 					phase++
// 				}
// 			}
// 		}
// 		phase++
// 		strs = nextStrs
// 	}
// }

// // 判断是否为回文串
// func checkPlalindrome(str string) bool {
// 	length := len(str)
// 	for i := 0; i < length/2; i++ {
// 		if str[length-1-i] != str[i] {
// 			return false
// 		}
// 	}
// 	return true
// }
