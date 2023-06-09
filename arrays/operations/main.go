package main

import (
	"fmt"
	"strings"
)

func main() {
	numbers := [4]int{10, 20, 30}

	fmt.Printf("%#v\n", numbers)

	numbers[0] = 7
	fmt.Printf("%#v\n", numbers)

	// numbers[5] = 100 // index out of bounds error

	for i, v := range numbers {
		fmt.Println("index:", i, "value:", v)
	}

	fmt.Println(strings.Repeat("#", 20))

	for i := 0; i < len(numbers); i++ {
		fmt.Println("index:", i, "value:", numbers[i])
	}

	//multidimensional array
	balances := [2][3]int{
		{5, 5, 7},
		[3]int{8,9,10},
	}
	fmt.Printf("%#v\n", balances)


    //  = operator creates a copy of an array.
    // the arrays are not connected and are saved in different memory locations
    m := [3]int{1, 2, 3}
    n := m //n is a copy of m
 
    fmt.Println("n is equal to m: ", n == m) // => true
    m[0] = -1                                //only m is modified
    fmt.Println("n is equal to m: ", n == m) // => false
}