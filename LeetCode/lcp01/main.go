package main

import "fmt"

func main() {
	fmt.Println(game([]int{2, 2, 3}, []int{3, 2, 1}))
}

func game(guess []int, answer []int) int {
	var (
		index int
		count int
	)
	for len(guess) == len(answer) && index < len(guess) {
		if guess[index] == answer[index] {
			count++
		}
		index++
	}
	return count
}
