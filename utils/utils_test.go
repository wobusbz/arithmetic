package utils

import (
	"arithmetic/test2"
	"testing"
)

func TestCalcRunTime(t *testing.T) {
	intArr := RandList(100000)
	CalcRunTime1("BubbleSort", test2.BubbleSort, intArr)
}

func TestRandList(t *testing.T) {

}
