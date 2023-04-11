package main

import (
	"fmt"
)

func main() {
	price, inStock := 100, true

	if price > 80 {
		fmt.Println("Too Expensive!")
	}

	if price <= 100 && inStock {
		fmt.Println("Buy it!")
	}

	// in go there's no such thing as truthy/falsy variable like javascript
	// if price { // not valid in go
	// 	fmt.Println("Something")
	// }

	if price < 100 {
		fmt.Println("It's cheap")
	} else if price == 100 {
		fmt.Println("On the edge")
	} else {
		fmt.Println("It's Expensive!")
	}

	age := 28

	if age >= 0 && age < 18 {
		fmt.Printf("You cannot vote! Please return in %d years!\n", 18 - age)
	} else if age == 18 {
		fmt.Printf("This is your first vote!\n")
	} else if age > 18 && age <= 100 {
		fmt.Printf("Please vote, it's important!!\n")
	} else {
		fmt.Println("Invalid age")
	}
}