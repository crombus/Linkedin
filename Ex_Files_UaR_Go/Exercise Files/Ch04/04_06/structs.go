package main

import (
	"fmt"
)

type dog struct {
	breed  string
	weight int
}

func main() {
	poddle := dog{"vella", 2345}
	fmt.Println(poddle)
	fmt.Printf("%+v\n", poddle)
	fmt.Printf("breed:%v\nweight:%v\n", poddle.breed, poddle.weight)

}
