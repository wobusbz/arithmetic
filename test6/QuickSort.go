package test6

// 快速排序：
// 需要找到基点
func QuickSort(intArr []int, startIndex, endIndex int) {
	if startIndex >= endIndex {
		return
	}
	midIndex := partition(intArr, startIndex, endIndex)
	QuickSort(intArr, startIndex, midIndex-1)
	QuickSort(intArr, midIndex+1, endIndex)
}

func partition(intArr []int, startIndex int, endIndex int) int {
	var pivot = intArr[startIndex]
	var left = startIndex
	var right = endIndex
	for left != right {
		for left < right && intArr[right] > pivot {
			right--
		}
		for left < right && intArr[left] <= pivot {
			left++
		}
		if left < right {
			intArr[left], intArr[right] = intArr[right], intArr[left]
		}
	}
	intArr[startIndex], intArr[left] = intArr[left], pivot
	return left
}
