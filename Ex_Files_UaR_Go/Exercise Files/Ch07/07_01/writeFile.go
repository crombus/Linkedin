package main

import (
	"fmt"
	"io"

	"os"
)

func main() {

	content := "Hello from Go!"
	file, err := os.Create("./tostring.txt")
	checkError(err)

	defer file.Close()
	ln, err := io.WriteString(file, content)
	checkError(err)

	fmt.Println("wrote string of ", ln, " characters")

	//bites := []byte(content)
	//ioutil.WriteFile("./tostring.txt", bites, 666)

}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

}
