package _5

func SeachArr(intArr []int, val int) int {
	var i int
	for i < len(intArr)-1 && intArr[i] <= val {
		if intArr[i] == val {
			return i
		}
		// 走到这里说明还没有找到，则把下标+1，继续向下找
		i++
	}
	return i
}
