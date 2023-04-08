package main

import (
	"fmt"
)

func main() {
	// unused constants are allowed but not variables
	const days int = 23

	const pi float64 = 3.14
	const secondsInHour = 3600

	duration := 234 // in hours
	fmt.Printf("Duration in seconds: %v\n", duration*secondsInHour)

	// variables belong to runtime
	x, y := 5, 0
	_, _ = x, y

	// fmt.Println(x / y) // runtime error

	// constants belong to compile time
	const a = 5
	const b = 0

	// fmt.Println(a / b) // compile time error

	const n, m int = 5, 7 // declaring with type
	const n1, m1 = 6, 9 // declaring without type	

	// const (
	// 	min1 = -500
	// 	min2 = -300
	// 	min3 = 100
	// )

	// in grouped constants the constants repeat the previous one
	const (
		min1 = -500
		min2
		min3
	)

	fmt.Println(min1, min2, min3) // out: -500 -500 -500

	
	// Constant Rules:
	// 1. You can not change a constant
	const temp int = 50
	// temp = 65
	
	// 2. You can not initiate a constant at runtime
	// const power = math.Pow(2, 3)
	
	// 3. You can not use a variable to initialize a constant
	// t := 5
	// const tc = t

	// 4. You can use func like len to initialize a constant
	// because string literal is an unname constant, len is builtin func available at compile time
	// but math.Pow is runtime func
	const l1 = len("hello")
}