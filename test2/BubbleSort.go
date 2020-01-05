package test2


func BubbleSort(intArr []int) {
	l := len(intArr)
	for i := 0; i < l; i++ { //  外面这个是控制循环查的次数
		var flag bool
		for j := i; j < l-i-1; j++ {
			if intArr[j] > intArr[j+1] {
				intArr[j], intArr[j+1] = intArr[j+1], intArr[j]
				flag = true
			}
		}
		if !flag {
			return
		}
	}
}
