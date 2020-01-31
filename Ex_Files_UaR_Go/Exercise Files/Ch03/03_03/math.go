package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {

	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("Integer sum: ", intSum)

	f1, f2, f3 := 12.5, 45.6, 65.7
	fsum := f1 + f2 + f3
	fmt.Printf("sum = %.2f\n", fsum)
	//if i write %.20f output = sum = 123.80000000000001136868

	var v1, v2, v3, bigsum big.Float
	v1.SetFloat64(23.5)
	v2.SetFloat64(34.5)
	v3.SetFloat64(45.6)
	bigsum.Add(&v1, &v2).Add(&bigsum, &v3)
	fmt.Printf("sum is =%.010g\n", &bigsum)

	cr := 24
	cicum := float64(cr) * math.Pi
	fmt.Println(cicum)
}
