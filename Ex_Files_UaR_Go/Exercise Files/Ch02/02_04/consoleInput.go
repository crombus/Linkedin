package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var s string
	fmt.Scanln(&s)
	fmt.Println(s)
	//fails when input "wer try iul"
	//output "wer"

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter text: ")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)
	//problem solved input enter "pulkit is good"
	// output pulkit is good

	fmt.Println("enter number: ")
	str, _ = reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	// if not use trimspace err: hello strconv.ParseFloat: parsing "34\n": invalid syntax
	if err != nil {
		fmt.Println("hello", err)
	} else {
		fmt.Println(f)
	}
}
