package main

import (
	"fmt"
)

func main() {
	var colors [3]string
	colors[0] = "red"
	colors[1] = "green"
	colors[2] = "yello"
	fmt.Println(colors)
	fmt.Println(colors[0])

	var num = [4]int{1, 2, 3, 4}
	fmt.Println(num)

	fmt.Println(len(colors[1]))
	fmt.Println(len(num))

	var co = [2]int{1, 2}
	fmt.Println(co)
}
