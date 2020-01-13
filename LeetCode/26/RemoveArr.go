package _6

func RemoveArr(nums []int) int {
	length := len(nums)
	k := 0
	for i := 1; i < length; i++ {
		if nums[k] != nums[i] {
			k++
			nums[k], nums[i] = nums[i], nums[k]
		}
	}
	return k + 1
}
