package mysort

func Bubble(slice []int) []int {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				tmp := slice[j]
				slice[j] = slice[j+1]
				slice[j+1] = tmp
			}
		}
	}
	return slice
}
