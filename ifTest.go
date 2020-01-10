package main

func main() {
	x := 10
	println(x == 10)
	if x < 5 {
		println("small")
	} else if x < 12 {
		println("medium")
	} else {
		println("big")
	}

	y := 12
	if x < 20 && y >= 20 {
		println("good")
	} else if x < 20 || y >= 20 {
		println("bad")
	}

	if xy := x + y; true {
		println(xy)
	}

	if y = 20; true {
		println(y)
	}

	if z, ok := y-x, true; ok {
		println(z)
	}
}
