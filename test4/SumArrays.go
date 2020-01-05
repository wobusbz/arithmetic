package test4

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

给定 nums = [2, 7, 11, 15], target = 9
9-2
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/

func SumArrays(intArr []int, target int) {
	var i int
	var index int
	var l = len(intArr) - 1
	for {
		if index > l {
			index = 0
			i++
		}
		temp := target - intArr[index]
		if intArr[i] == temp {
			return
		}
		//if intArr[i]+intArr[index] == target {
		//	//fmt.Printf("%s 下标为: %d %d\n", "找到了", i, index)
		//	//fmt.Printf("%d + %d = %d\n", intArr[i], intArr[index], intArr[i]+intArr[index])
		//	return
		//}
		//fmt.Printf("%d + %d = %d\n", intArr[i], intArr[index], intArr[i]+intArr[index])
		index++
	}
}

func SumArrays1(intArr []int, target int) {
	var mapInt = make(map[int]int, len(intArr))
	for k, v := range intArr {
		mapInt[v] = k
	}
	for _, v := range mapInt {
		temp := target - v
		if _, ok := mapInt[temp]; ok {
			return
		}
	}
}
