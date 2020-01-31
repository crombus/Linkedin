package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("filename.ext")
	if err == nil {
		fmt.Println("hello", f)
	} else {
		fmt.Println(err.Error())
	}
	myerror := errors.New("hey mojaboa")
	fmt.Println(myerror)

	// attendance := map[string]bool{
	// 	"Ann": true,
	// 	"Mike": true}
}
