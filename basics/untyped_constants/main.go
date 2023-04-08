package main

import (
	"fmt"
)

func main() {
	const a float64 = 5.1 // typed constant
	const b = 6.7 // untyped constant

	const c float64 = a * b
	const str = "Hello " + "Go!"

	const d =  5 > 10
	fmt.Println(d)

	// const x int = 5 // typed constant
	// const y float64 = 2.2 * x // not allowed to multiply int const with untyped constant 2.2
	const x = 5 // untyped constant
	const y = x * 2.2 // multiplication of untyped constants are allowed
	fmt.Printf("%T\n", y)

	var i int = x // x changes to int
	var j float64 = x // var j float64 = float64(x) // go does this behind the scenes
	var  p byte = x
	fmt.Println(i, j, p)
}