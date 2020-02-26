package main

import	(
	"fmt"
	"encoding/base64"
)


func main() {
	name := "delrt"
	ep := "rty"
	erl := fmt.Sprintf("{Name : %s , Endpoint : %s}", name , ep)
	fmt.Println(erl)
        pwd := ("{Name : "+name+" , Endpoint : "+ep+"}")	
	encoded := base64.StdEncoding.EncodeToString([]byte(pwd))
	fmt.Println(encoded)
}
