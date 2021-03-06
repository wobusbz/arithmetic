package rotateList

func RotateList(intArr []int32) []int32 {
	if len(intArr) < 1 {
		return intArr
	}
	for i := 0; i < len(intArr)/2; i++ {
		intArr[i], intArr[len(intArr)-1-i] = intArr[len(intArr)-1-i], intArr[i]
	}
	return intArr
}

func Roate(intArr []int32, n int) {
	RotateList(intArr)
	RotateList(intArr[:2%len(intArr)])
	RotateList(intArr[2%len(intArr):])
}
