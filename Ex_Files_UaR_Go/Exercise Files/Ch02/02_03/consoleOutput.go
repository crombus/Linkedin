package main

import "fmt"

func main() {

	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	aNumber := 42
	isTrue := true
	strl, _ := fmt.Println(str1, str2, str3)

	fmt.Println("string length:", strl)

	fmt.Printf("this yo %v\n", aNumber)
	fmt.Printf("this is wo %v\n", isTrue)

	fmt.Printf("this is float %08.2f\n", float64(aNumber))
	//output this is float 00042.00

	fmt.Printf("this is type %T %T %T", str3, aNumber, isTrue)
	// output this is type string int bool

}
