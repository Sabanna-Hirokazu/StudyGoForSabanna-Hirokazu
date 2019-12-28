package main

import "fmt"

func remove(slice []int, location int) []int {
	if location < 0 || location >= len(slice) {
		fmt.Println("Error: invalid index number")
		return slice
	}
	result := slice[:location]
	result = append(result, slice[location+1:]...)
	return result
}

func main() {
	array := [...]int{10, 11, 12}
	slice := []int{}

	fmt.Println(array)
	fmt.Println(slice)

	slice = append(slice, 5)
	slice = append(slice, 7)
	slice = append(slice, 9)
	fmt.Println(slice)

	array[1] += 9
	slice[2] = 15

	fmt.Println(array)
	fmt.Println(slice)

	slice = remove(slice, 1)
	fmt.Println(slice)
}
