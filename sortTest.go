package main

import (
	"fmt"

	"./mysort"
)

func main() {
	tmp1 := []int{3, 5, 2, 4, 1}
	tmp2 := []int{4, 2, 6, 8, 1, 3, 7, 5}
	fmt.Println(tmp1)
	fmt.Println(tmp1[:3])
	fmt.Println(tmp1[3:])

	mysort.Help()

	fmt.Println(mysort.Bubble(tmp1))
	tmp1 = []int{3, 5, 2, 4, 1}
	fmt.Println(mysort.Selection(tmp1))
	tmp1 = []int{3, 5, 2, 4, 1}
	fmt.Println(mysort.Insert(tmp1))
	tmp1 = []int{3, 5, 2, 4, 1}
	fmt.Println(mysort.Merge(tmp1))
	tmp1 = []int{3, 5, 2, 4, 1}
	fmt.Println(mysort.Quick(tmp1))
	fmt.Println(mysort.Quick(tmp2))
}
