package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(InterReversal(-10))
	fmt.Println(math.MinInt32)
}

func InterReversal(x int) (num int) {
	if num == -1 {
		return 0
	}
	for x != 0 {
		num = num*10 + x%10
		x = x / 10
	}
	if num > 1<<31-1 || num < -1<<31-1 {
		return 0
	}
	return num
}
