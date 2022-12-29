package solution120

func maxProfit1(prices []int) int {
	// DP解法：用一个数组来分别表示持有股票和不持有股票时的最大收益
	// 最后一位不持有股票时的最大收益就是解
	dp := make([][]int, len(prices))
	dp[0] = []int{0, -prices[0]}
	for i := 1; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[len(dp)-1][0]
}
