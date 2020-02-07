package mysort

func selectPivot(slice []int) (int, int) {
	// 次の値を見て先頭か先頭より大きい値を選択する
	if slice[0] < slice[1] {
		return slice[1], 1
	}
	return slice[0], 0
}

func separateForPivot(slice []int) []int {
	if len(slice) <= 1 {
		return slice
	}
	pivot, left := selectPivot(slice)
	right := len(slice) - 1
	for left < right {
		for pivot > slice[left] {
			left++
		}
		for pivot < slice[right] {
			right--
		}
		if left >= right {
			break
		}
		tmp := slice[left]
		slice[left] = slice[right]
		slice[right] = tmp
	}
	return append(separateForPivot(slice[:left]), separateForPivot(slice[left:])...)
}

func Quick(slice []int) []int {
	return separateForPivot(slice)
}
