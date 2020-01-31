package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(t)

	now := time.Now()
	fmt.Println(now)

	fmt.Println(t.Month(), t.Day(), t.Weekday())

	tom := t.AddDate(0, 0, 2)
	fmt.Println(tom.Weekday(), tom.Month(), tom.Day(), tom.Year())

	lf := "Monday,january 2,2006"
	fmt.Println(tom.Format(lf))
	sf := "2/1/06"
	fmt.Println(tom.Format(sf))

	fmt.Println(time.Minute)
}
