package test7

// 0 1 2 3 4 5
// 0 1 1 2 3 5 8 13 21
func FibonacciNumber(n int) int {
	if n <= 1 {
		return n
	}
	var first int
	var second = 1
	for i := 0; i < n-1; i++ {
		var sum = first + second // 0 + 1
		//fmt.Printf("%d = %d + %d\n", sum, first, second)
		first, second = second, sum //  前面的一个元素和sum交换  + sum 0+1 = 1 1+1 =2 2 +1 = 3
	}
	return second
}

func FibonacciNumber1(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciNumber(n-1) + FibonacciNumber(n-2)
}
