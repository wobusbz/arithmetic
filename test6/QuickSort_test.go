package test6

import (
	"arithmetic/utils"
	"testing"
)

func TestQuickSort(t *testing.T) {
	//intArr := []int{10000, 199,12, 5, 7,22, 6, 100, 99}
	//			   [7 22 12 5 22 199 10000 100 99]
	//intArr := []int{5, 7, 4, 6, 3, 1, 2, 9, 8}
	list := utils.RandList(100000000)
	l := len(list) - 1
	t.Log(l)
	//QuickSort(intArr, 0, len(intArr)-1)
	utils.CalcRunTime("QuickSort", QuickSort, list, 0, l)

	//t.Log(list)
}
