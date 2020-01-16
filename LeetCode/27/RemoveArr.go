package _7

import "fmt"

func RemoveArr(intArr []int, val int) int {

	var i, j = 0, len(intArr) - 1
	for {
		for i < len(intArr) && intArr[i] != val {
			i++
			fmt.Println("i:", i)
		}
		for j >= 0 && intArr[j] == val {
			j--
			fmt.Println("j:", j)
		}
		if i > j {
			break
		}
		intArr[i], intArr[j] = intArr[j], intArr[i]
	}
	return i
}
