package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	fileName := "./hello.txt"

	f, err := ioutil.ReadFile(fileName)
	checkError(err)
	fmt.Println(f)
	fmt.Println(string(f))

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
