package solution120

func maxProfit(prices []int) int {
	dp := make([]int, len(prices))
	dp[0] = 0
	min := prices[0]
	ans := 0
	for i := 1; i < len(prices); i++ {
		dp[i] = prices[i] - min
		if dp[i] < dp[i-1] {
			dp[i] = dp[i-1]
		}
		if prices[i] < min {
			min = prices[i]
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}
