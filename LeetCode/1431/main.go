package main

import "fmt"

func main() {
	fmt.Println(kidsWithCandies([]int{4,2,1,1,2}, 1))
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var maxNumber int
	for i := range candies {
		maxNumber = Max(maxNumber, candies[i])
	}
	var result = make([]bool, len(candies))
	for i := range candies {
		result[i] = candies[i] + extraCandies >= maxNumber
	}
	fmt.Println(maxNumber)
	return result
}

func Max(m, n int) int {
	if m > n {
		return m
	}
	return n
}