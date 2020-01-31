package main

import "fmt"

func main() {
	n1, n2 := fullname("john", "dick")
	fmt.Println(n1, n2)
	n1, n2 = fullnamenaked("john", "dick")
	fmt.Println(n1, n2)

}

func fullname(f, l string) (string, int) {
	full := f + " " + l
	le := len(f) + len(l)
	return full, le
}

func fullnamenaked(f, l string) (full string, le int) {
	full = f + " " + l
	le = len(f) + len(l)
	return
}
