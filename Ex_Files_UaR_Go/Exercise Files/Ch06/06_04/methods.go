package main

import "fmt"

type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

func (d Dog) speak() {
	fmt.Println(d.Sound)
}

func main() {
	po := Dog{"gty", 45, "woof"}
	po.speak()

}
