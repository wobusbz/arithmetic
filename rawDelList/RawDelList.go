package rawDelList

func RawDelList(list []int, val int) int {
	for i := 0; i < len(list); {
		if list[i] == val {
			list = append(list[:i], list[i+1:]...)
		} else {
			i++
		}
	}
	return len(list)
}
