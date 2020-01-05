package test5

import (
	"fmt"
	"time"
)

func InsertSort(intArr []int) {
	fmt.Println(intArr)
	for i := 1; i < len(intArr); i++ {
		for j := i; j > 0 && intArr[j-1] > intArr[j]; j-- {
			intArr[j], intArr[j-1] = intArr[j-1], intArr[j]
		}
		fmt.Printf("%d 次排序： %d\n", i, intArr)
	}
}

func InsertSort1(intArr []int) {
	fmt.Printf("原始数组: %v\n", intArr)
	var flag bool
	// 1 n
	// 1 n
	//O(n^2)
	for i := 0; i < len(intArr)-1; i++ { //
		for j := i + 1; j < len(intArr) && intArr[j-1] > intArr[j]; j++ {
			intArr[j], intArr[j-1] = intArr[j-1], intArr[j]
			flag = true
			time.Sleep(time.Second)
		}
		if !flag { // 这是极端的情况
			return
		}
		fmt.Printf("%d 次排序： %d\n", i, intArr)
	}
}
