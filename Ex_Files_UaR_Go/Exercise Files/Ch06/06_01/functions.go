package main

import "fmt"

func main() {

	sum := addsome(12, 34, 56, 78)
	fmt.Println(sum)
	fmt.Println(add(12, 34))
}
func addsome(valeues ...int) int {
	sum := 0
	for i := range valeues {
		sum += valeues[i]
	}
	fmt.Printf("%T\n", valeues)
	return sum
}
func add(x int, v int) int {
	//this will work also x, v int
	return x + v
}
