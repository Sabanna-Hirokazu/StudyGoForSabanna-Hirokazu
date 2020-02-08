package main

type Animal interface {
	Talk()
	Name()
}

type Dog struct {
	name string
}

func (d Dog) Talk() {
	println("Bow-wow!")
}

func (d Dog) Name() {
	println(d.name)
}

type Cat struct {
	name string
}

func (c Cat) Talk() {
	println("Meow!")
}

func (c Cat) Name() {
	println(c.name)
}

func main() {
	var animal Animal
	animal = Dog{name: "Taro"}
	animal.Talk()
	animal.Name()
	animal = Cat{name: "Jiro"}
	animal.Talk()
	animal.Name()
}
