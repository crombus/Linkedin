package main

import "fmt"

func main() {
	var p *int //this is nil pointer
	x := 10
	p = &x
	fmt.Println(*p)

	st := "tuyu"
	p1 := &st
	fmt.Println(*p1)

	*p = *p / 2
	fmt.Println(*p, x)
}
