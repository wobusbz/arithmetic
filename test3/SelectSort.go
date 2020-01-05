package test3

import (
	"fmt"
	"time"
)

func SelectSort(intArr []int) {
	var minIndex int
	for i := 0; i < len(intArr)-1; i++ {
		minIndex = i
		for j := i + 1; j < len(intArr); j++ {
			if intArr[j] < intArr[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i{
			intArr[i], intArr[minIndex] = intArr[minIndex], intArr[i]
		}
		fmt.Printf("minIndex: %d->%d i: %d->%d\n", minIndex, intArr[minIndex], i, intArr[i])
		fmt.Println(intArr)
		time.Sleep(time.Second)
	}
}
