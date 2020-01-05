package test1

import "fmt"

func BinarySearch(intArr []int, i int) {
	lMin := 0
	lMax := len(intArr) - 1
	for lMin <= lMax {
		mid := lMin + (lMax-lMin)/2
		fmt.Println("mid index: ", mid)
		midValue := intArr[mid]
		if midValue < i {
			lMin = mid + 1
			fmt.Println("lMin: ", midValue)
		} else if midValue > i {
			lMax = mid - 1
			fmt.Println("lMax: ", midValue)
		} else if midValue == i {
			fmt.Println("找到了", midValue)
			return
		}
	}
	fmt.Println("没有找到")
}
