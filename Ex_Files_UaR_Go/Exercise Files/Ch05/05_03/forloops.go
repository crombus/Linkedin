package main

import "fmt"

func main() {

	sum := 1
	fmt.Println("Sum:", sum)

	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	for i := 0; sum < 5; i++ {
		sum += i
		fmt.Println(sum)
	}
	for i := range colors {
		fmt.Println(i)
		goto gotoprogram
	}
gotoprogram:
	fmt.Println("haha")
}
