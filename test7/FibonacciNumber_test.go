package test7

import (
	"arithmetic/utils"
	"testing"
)

func TestFibonacciNumber(t *testing.T) {
	t.Log(FibonacciNumber(100))
	utils.CalcRunTime2("FibonacciNumber", FibonacciNumber, 10)
}

func TestFibonacciNumber1(t *testing.T) {
	t.Log(FibonacciNumber1(100))
	utils.CalcRunTime2("FibonacciNumber", FibonacciNumber1, 64)
}
