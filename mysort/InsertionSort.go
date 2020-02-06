package mysort

func Insert(slice []int) []int {
	for i := 1; i < len(slice); i++ {
		var j int
		tmp := slice[i]
		for j = i - 1; j >= 0; j-- {
			if tmp < slice[j] {
				slice[j+1] = slice[j]
			} else {
				break
			}
		}
		slice[j+1] = tmp
	}
	return slice
}
