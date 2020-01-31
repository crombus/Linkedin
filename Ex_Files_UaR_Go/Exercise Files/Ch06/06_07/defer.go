package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")
	defer fmt.Println("hello late")
	fmt.Println("hellocaleed after lster")
	myFunc()
	fmt.Println("heldefefe")
	defer fmt.Println("hello later very kate")
	fmt.Println("simple")

}

func myFunc() {
	defer fmt.Println(":yello")
	fmt.Println("hello orfehf")
}
