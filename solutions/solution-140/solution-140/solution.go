package solution140

import (
	"strings"
)

func wordBreak(s string, wordDict []string) []string {
	wordMap := make(map[string]int)
	for i, word := range wordDict {
		wordMap[word] = i
	}

	dp := make([]bool, len(s)+1)
	dp2 := make([][][]string, len(s)+1)

	dp[0] = true

	for i := 1; i < len(dp); i++ {
		for j := 0; j < i; j++ {
			_, exist := wordMap[s[j:i]]
			if dp[j] && exist {
				dp[i] = true
				var list [][]string
				if dp2[j] == nil {
					list = [][]string{}
				} else {
					list = dp2[j]
				}

				if dp2[i] == nil {
					dp2[i] = [][]string{}
				}
				if len(list) == 0 {
					dp2[i] = append(dp2[i], []string{s[j:i]})
				} else {
					for _, v := range list {
						l := append([]string{}, v...)
						l = append(l, s[j:i])
						dp2[i] = append(dp2[i], l)
					}
				}
			}
		}
	}

	list := dp2[len(s)]
	ans := make([]string, len(list))
	for i, v := range list {
		ans[i] = strings.Join(v, " ")
	}

	return ans
}
