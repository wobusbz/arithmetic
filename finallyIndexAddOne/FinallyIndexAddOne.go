package finallyIndexAddOne

func FinallyIndexAddOne(intS []int) []int {
	var len = len(intS)
	var num = intS[len-1] + 1
	var newIntS []int
	if num > 10 { // 9 + 1 = 10
		newIntS = append(newIntS, num-10)
	} else {
		newIntS = append(newIntS, num)
	}

	for i := len - 2; i > 0; i-- {

	}
	return newIntS
}
