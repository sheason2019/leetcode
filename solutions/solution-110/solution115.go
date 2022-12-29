package solution110

func numDistinct(s string, t string) int {
	// 初始化DP
	dp := make([][]int, len(t)+1)
	for i := range dp {
		dp[i] = make([]int, len(s)+1)
	}
	for i := range dp[0] {
		dp[0][i] = 1
	}

	for i := 1; i < len(t)+1; i++ {
		for j := i; j < len(s)+1; j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	return dp[len(t)][len(s)]
}
