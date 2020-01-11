package main

import "fmt"

func main() {
	dic := map[string]int{"A": 1, "B": 2, "C": 3}
	number, ok := dic["A"]
	fmt.Println(number)
	fmt.Println(ok)
}
