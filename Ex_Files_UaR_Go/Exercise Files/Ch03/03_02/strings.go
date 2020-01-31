package main

import (
	"fmt"
	"strings"
)

func main() {
	val1 := "this implict dec"
	fmt.Println(val1)
	//var string string = "this explicit dec"
	//fmt.Println(string)
	//string is data type but work as variable also
	var val2 string = "this is also explicit dec"
	fmt.Println(val2)
	//if string used before then this line gives error because
	// string is already declared

	fmt.Println(strings.ToUpper(val1))
	fmt.Println(strings.Title(val2))

	exp1 := "hello"
	exp2 := "HELLO"

	fmt.Println(exp1 == exp2)
	fmt.Println(strings.EqualFold(exp1, exp2))

	fmt.Println(strings.Contains(val1, "exp"))
	fmt.Println(strings.Contains(val2, "exp"))

}
