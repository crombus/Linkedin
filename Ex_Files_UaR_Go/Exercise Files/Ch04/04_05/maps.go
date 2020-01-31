package main

import (
	"fmt"
	"sort"
)

func main() {
	states := make(map[string]string)
	fmt.Println(states)
	states["PkL"] = "panchkula"
	states["chd"] = "chandigarh"
	states["blr"] = "bangalore"
	fmt.Println(states)

	fmt.Println(states["chd"])

	delete(states, "chd")
	fmt.Println(states)
	states["hyd"] = "hyderabad"
	fmt.Println(states)

	for i, k := range states {
		fmt.Println(i + "-----" + k)
		fmt.Printf("%v-----%v", i, k)
	}

	k := 0
	keys := make([]string, len(states))
	for i := range states {
		keys[k] = i
		k++
	}
	sort.Strings(keys)
	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}
