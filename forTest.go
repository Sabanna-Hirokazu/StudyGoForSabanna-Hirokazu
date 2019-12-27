package main

import "strconv"

func main() {
	count := 1
	// 無限ループと等しい
	for {
		if count > 5 {
			println("break")
			break
		}
		println(count)
		count++

	}
	for count != 0 {
		println(count)
		count -= 2
	}

	count = 0
LABEL1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 0 && j == 2 {
				// iのforへcontinue
				println("continue Label")
				continue LABEL1
			} else if i == 1 && j == 1 {
				// jのforへcontinue
				println("continue")
				continue
			}
			count++
		}
	}

	println("count: " + strconv.Itoa(count))

Label2:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			println(i + j)
			continue Label2
		}
	}
}
