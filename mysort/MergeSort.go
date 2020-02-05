package mysort

func merge(left []int, right []int) []int {
	result := []int{}
	valLeft := len(left)
	valRight := len(right)
	i, j := 0, 0
	for i < valLeft || j < valRight {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	if i < valLeft {
		result = append(result, left[i:]...)
	} else {
		result = append(result, right[j:]...)
	}
	return result
}

func separate(slice []int) []int {
	if len(slice) < 1 {
		return slice
	}
	var middle int = len(slice) / 2
	return merge(separate(slice[:middle]), separate(slice[middle:]))
}

func Merge(slice []int) []int {
	return slice
}
