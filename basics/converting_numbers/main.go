package main

import (
	"fmt"
)

func main() {
	var x = 3 // type int
	var y = 3.1 // type float64

	// x = x * y
	x = x * int(y)

	fmt.Println(x)
	fmt.Printf("%T\n", y)

	x = int(float64(x) * y)
	fmt.Printf("%d\n", x)

	y = float64(x) * y
	fmt.Println(y)

	var a = 5 // type int
	fmt.Printf("%T\n", a)

	var b int64 = 2
	a = int(b)
}