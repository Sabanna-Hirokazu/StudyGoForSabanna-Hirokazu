package main

func main() {
	count := 5
	switch count % 4 {
	case 0:
		println(0)
	case 1:
		println(1)
	case 2:
		println(2)
	default:
		println(3)
	}

	switch count {
	// count == 2 || count == 3 と等しい
	case 2, 3:
		println(23)
	case 4, 5:
		println(45)
	default:
		println(0)
	}
}
