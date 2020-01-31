package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
	dow := rand.Intn(6) + 1
	fmt.Println("Day", dow)

	switch dow {
	case 1:
		fmt.Println("monday")
	case 5:
		fmt.Println("friday")
	default:
		fmt.Println("weekday")
	}

}
