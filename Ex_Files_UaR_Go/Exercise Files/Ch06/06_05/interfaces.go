package main

import "fmt"

type animal interface {
	speak() string
}
type Dog struct {
}

func (d Dog) speak() string {
	fmt.Println("wooof")
	return "hi"
}

func main() {
	poodle := animal(Dog{})
	fmt.Println(poodle)
	fmt.Println(poodle.speak())
}
