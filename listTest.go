package main

import (
	"fmt"
	"strconv"
)

func main() {
	a1 := [3]int{1, 2, 3}
	a2 := [...]int{4, 5, 6}
	fmt.Println(a1)
	fmt.Println(a2)
	println(a1[0])
	println("number: " + strconv.Itoa(a1[1]))
	println(a1[2])

	s1 := []int{}
	fmt.Println(s1)
	s1 = append(s1, 10)
	fmt.Println(s1)
	s1 = append(s1, 20)
	fmt.Println(s1)
	t := append(s1, 30)
	fmt.Println(s1)
	fmt.Println(t)

	s2 := [][]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(s2)
	s2[0] = append(s2[0], 9)
	fmt.Println(s2)

}
