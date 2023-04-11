package main

import (
	"fmt"
)

func main() {
	grades := [5]int{
		1: 10,
		0: 5,
		4: 7,
	}
	fmt.Println(grades)

	names := [...]string{
		5: "Dan",
	}
	fmt.Printf("%#v\n", names)

	cities := [...]string{
		5: "Paris",
		"London", // index 6
		1: "NYC",
	}
	fmt.Printf("%#v\n", cities)
}