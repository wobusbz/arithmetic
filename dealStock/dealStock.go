package dealStock

func DealStock(prices []int32) int32 {
	if len(prices) == 0 {
		return 0
	}
	var maxProfit int32

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}
	return maxProfit
}

func Max(min, max int32) int32 {
	if min > max {
		return min
	}
	return max
}
