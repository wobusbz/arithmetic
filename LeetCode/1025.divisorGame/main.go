package main

import "fmt"

func main() {
	fmt.Println(divisorGame(12))
}

func divisorGame(n int) bool {
	return n&1 == 1
}
