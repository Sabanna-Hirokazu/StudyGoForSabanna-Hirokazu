package main

import (
	"fmt"
	"strconv"
)

type human struct {
	name string
	age  int
}

func (h human) message() string {
	return h.name + "は" + strconv.Itoa(h.age) + "歳です。"
}

func (h *human) template() {
	h.name = "Hoge"
	h.age = 20
}

type japanese struct {
	blood   string
	element human
}

type japanese2 struct {
	blood   string
	element *human
}

func main() {

	h1 := new(human)
	fmt.Println(h1)
	h1.name = "Steve"
	h1.age = 24
	fmt.Println(h1)
	fmt.Println(h1.message())
	h1.template()
	fmt.Println(h1)

	h2 := human{name: "Smith", age: 45}
	fmt.Println(h2)
	h2.name += ".A"
	h2.age += 5
	fmt.Println(h2)
	fmt.Println(h2.message())
	h2.template()
	fmt.Println(h2)

	h3 := japanese{blood: "A", element: h2}
	fmt.Println(h3)
	fmt.Println(h3.element)

	h4 := japanese2{blood: "A", element: h1}
	fmt.Println(h4)
	fmt.Println(h4.element)
}
