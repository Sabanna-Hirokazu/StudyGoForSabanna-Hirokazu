package main

import "fmt"

func main() {
	var n1 int = 10
	var n2 *int = &n1
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(*n2)

	n1 += 15
	fmt.Println(n1)
	fmt.Println(*n2)
}
