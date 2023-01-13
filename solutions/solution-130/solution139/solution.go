package solution139

func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	// 构建Set
	for _, word := range wordDict {
		wordSet[word] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 0; i < len(dp); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
			}
		}
	}

	return dp[len(s)]
}
