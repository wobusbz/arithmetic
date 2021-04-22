package main

import "fmt"

func main() {
	fmt.Println(xorOperation(4, 3))
	fmt.Println(1<<0 + 1<<1)
}

func xorOperation(n int, start int) int {
	var result int
	for i := 0; i < n; i++ {
		result ^= start + 2*i
	}
	return result
}

// 0000 0000 0000 0011    
// 0000 0000 0000 0101
// 0000 0000 0000 0110
