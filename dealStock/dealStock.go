package dealStock

import (
	"fmt"
)

func DealStock(prices []int32) int32 {
	if len(prices) < 2 {
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

func DealStock1(prices []int32) int32 {
	if len(prices) < 2 {
		return 0
	}

	var maxPrices = make([][2]int32, len(prices))
	maxPrices[0][0] = 0
	maxPrices[0][1] = -prices[0]

	for i := 1; i < len(prices); i++ {
		maxPrices[i][0] = Max(maxPrices[i-1][0], maxPrices[i-1][1]+prices[i])
		maxPrices[i][1] = Max(maxPrices[i-1][0]-prices[i], maxPrices[i-1][1])
	}
	fmt.Println(maxPrices)
	return maxPrices[len(prices)-1][0]
}

func Max(min, max int32) int32 {
	if min > max {
		return min
	}
	return max
}
