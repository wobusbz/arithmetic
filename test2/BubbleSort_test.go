package test2

import (
	"arithmetic/utils"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	intArr := []int{10, 2, 5, 7, 6, 100, 99, 33}
	BubbleSort(intArr)
	list := utils.RandList(100000)
	utils.CalcRunTime1("BubbleSort", BubbleSort, list)
	t.Log(intArr)
}
