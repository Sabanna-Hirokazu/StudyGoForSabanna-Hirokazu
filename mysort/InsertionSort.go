package mysort

func Insert(slice []int) []int {
	for i := 1; i < len(slice); i++ {
		var j int
		for j = i - 1; j > 0; j-- {
			if slice[i] < slice[j] {
				slice[j+1] = slice[j]
			} else {
				break
			}
		}
		slice[j+1] = slice[i]
	}
	return slice
}
