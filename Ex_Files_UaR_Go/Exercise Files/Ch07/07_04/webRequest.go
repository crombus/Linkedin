package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "http://services.explorecalifornia.org/json/tours.php"
	resp, _ := http.Get(url)

	f, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(f))
	resp.Body.Close()

}
