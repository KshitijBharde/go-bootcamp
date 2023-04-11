package main

import (
	"fmt"
	"strconv"
	"os"
)

func main() {
	i, err := strconv.Atoi("45")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

	// simple (short) statement ->  the same effect as the above code
    // i and err are variables scoped to the if statement only
	if i, err := strconv.Atoi("26a"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

	if args := os.Args; len(args) < 2 {
		fmt.Println("At leaset on argument is required")
	} else if km, err := strconv.Atoi(os.Args[1]); err != nil {
		fmt.Println("The argument must be an integer, Error:", err)
	} else {
		fmt.Printf("%d kilometers in miles is %v\n", km, float64(km)*0.621)
	}
}