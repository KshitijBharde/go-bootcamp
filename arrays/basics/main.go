package main

import (
	"fmt"
)

func main() {
	var numbers [4]int

	fmt.Printf("%v\n", numbers)
	fmt.Printf("%#v\n", numbers)

	var arr1 = [4]float64{-10, 1, 100}
	fmt.Printf("%v\n", arr1)

	arr2 := [4]string{"a", "b", "c"}
	fmt.Printf("%#v\n", arr2)

	// ellipsis operator
	arr3 := [...]int{
		21,
		2423,
		54,
		545,
		65,
	}
	fmt.Printf("%#v\n", arr3)
}