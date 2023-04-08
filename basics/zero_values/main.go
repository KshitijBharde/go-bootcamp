package main

import "fmt"

func main() {
	// type inference
	var a = 4
	var b = 10.8

	// a = b // gives error: cannot use b (variable of type float64) as int value in assignment 

	a = int(b)
	fmt.Println(a, b)

	/*
	In go all variable are initialized
	- numeric type: 0
	- bool type: false
	- string type: ""
	- pointer type: nil
	*/
	var value int
	var price float64
	var name string
	var done bool
	fmt.Println(value, price, name, done)
}