package test7

import (
	"arithmetic/utils"
	"testing"
)

func TestFibonacciNumber(t *testing.T) {
	t.Log(FibonacciNumber(10))
	utils.CalcRunTime2("FibonacciNumber", FibonacciNumber, 10)
}

func TestFibonacciNumber1(t *testing.T) {
	utils.CalcRunTime2("FibonacciNumber", FibonacciNumber1, 10)
}
