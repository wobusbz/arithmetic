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

func Quicksort(array []int, begin, end int) {
	var i, j int
	if begin < end {
		i = begin + 1 // 将array[begin]作为基准数，因此从array[begin+1]开始与基准数比较！
		j = end       // array[end]是数组的最后一位

		for {
			if i >= j {
				break
			}
			if array[i] > array[begin] {
				array[i], array[j] = array[j], array[i]
				j = j - 1
			} else {
				i = i + 1
			}

		}
		if array[i] >= array[begin] { // 这里必须要取等“>=”，否则数组元素由相同的值时，会出现错误！
			i = i - 1
		}
		array[begin], array[i] = array[i], array[begin]
		Quicksort(array, begin, i)
		Quicksort(array, j, end)
	}
}
