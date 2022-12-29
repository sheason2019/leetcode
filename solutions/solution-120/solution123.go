package solution120

func maxProfit2(prices []int) int {
	buy1 := -prices[0]
	sell1 := 0
	buy2 := -prices[0]
	sell2 := 0

	for _, price := range prices {
		buy1 = max(buy1, -price)
		sell1 = max(sell1, price+buy1)
		buy2 = max(buy2, sell1-price)
		sell2 = max(sell2, price+buy2)
	}

	return sell2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
