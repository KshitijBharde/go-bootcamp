package main

import (
	"fmt"
)

func main() {
	var n []int
	fmt.Println(n == nil)
	fmt.Printf("%#v\n", n)

	m := []int{}
	fmt.Println(m == nil)
	fmt.Printf("%#v\n", m)

	a, b := []int{1, 2, 3}, []int{1, 2, 5}

	// fmt.Println(a == b) // gives error, can only compare slice using == operator with nil

	var eq bool = true

	for idx, valA := range a{
		if valA != b[idx] {
			eq = false
			break
		}
	}

	if len(a) != len(b) {
		eq = false
	}

	if eq {
		fmt.Println("a and b slices are equal")
	} else {
		fmt.Println("a and b slices are not equal")
	}

}