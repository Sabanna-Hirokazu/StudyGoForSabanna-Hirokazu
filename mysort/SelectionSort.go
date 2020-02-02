package mysort

func Selection(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		min := i
		for j := i + 1; j < len(slice); j++ {
			if slice[min] > slice[j] {
				min = j
			}
		}
		tmp := slice[i]
		slice[i] = slice[min]
		slice[min] = tmp
	}
	return slice
}
