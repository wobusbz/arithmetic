package test

import "fmt"

/*
	判断符号{}()的有效性
*/

func IsJudegeSymbol(symbol string) {
	fmt.Println(symbol)
	var re []byte
	var count int
	// O(n^2)
	for _, chr := range symbol {
		if chr == '(' || chr == ')' {
			re = append(re, byte(chr))
		}
	}
	//O(n)
	for i := 0; i < len(re)-1; i++ {
		if re[i] == '(' && re[i+1] == ')' {
			count++
		}
	}
	fmt.Println("symbol的有效性: ", count)
}


/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

给定 nums = [2, 7, 11, 15], target = 9
9-2
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
给定在 xy 平面上的一组点，确定由这些点组成的任何矩形的最小面积，其中矩形的边不一定平行于 x 轴和 y 轴。

如果没有任何矩形，就返回 0。



示例 1：



输入：[[1,2],[2,1],[1,0],[0,1]]
输出：2.00000
解释：最小面积的矩形出现在 [1,2],[2,1],[1,0],[0,1] 处，面积为 2。
示例 2：



输入：[[0,1],[2,1],[1,1],[1,0],[2,0]]
输出：1.00000
解释：最小面积的矩形出现在 [1,0],[1,1],[2,1],[2,0] 处，面积为 1。
示例 3：



输入：[[0,3],[1,2],[3,1],[1,3],[2,1]]
输出：0
解释：没法从这些点中组成任何矩形。
示例 4：



输入：[[3,1],[1,1],[0,1],[2,1],[3,3],[3,2],[0,2],[2,3]]
输出：2.00000
解释：最小面积的矩形出现在 [2,1],[2,3],[3,3],[3,1] 处，面积为 2。


提示：

1 <= points.length <= 50
0 <= points[i][0] <= 40000
0 <= points[i][1] <= 40000
所有的点都是不同的。
与真实值误差不超过 10^-5 的答案将视为正确结果
*/