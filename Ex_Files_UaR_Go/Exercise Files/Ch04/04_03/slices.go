package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"red", "der", "frty"}
	fmt.Println(colors)
	fmt.Println(append(colors, "purple"))

	var nu = make([]int, 5, 5)
	nu[1] = 23
	nu[0] = 456
	nu[2] = 1
	nu[3] = 345
	fmt.Println(nu)
	sort.Ints(nu)
	fmt.Println(nu)

}
