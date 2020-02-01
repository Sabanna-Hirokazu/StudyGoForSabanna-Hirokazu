package main

import (
	"fmt"

	"./mysort"
)

func main() {
	// hello, world
	mysort.PrintTest()

	tmp1 := []int{3, 5, 2, 4, 1}
	fmt.Println(mysort.Bubble(tmp1))
}
