package main

import "fmt"

func main() {
	var age int = 28
	fmt.Println("Age:", age)

	var name = "Kshitij"
	// fmt.Println("Your name is: ", name)
	_ = name // unused variable using Blank Identifier

	// use short declaration when we know the initial value
	s := "Learning golang"
	fmt.Println(s)

	car, cost := "Audi", 50000
	fmt.Println(car, cost)

	car, year := "BMW", 2023
	fmt.Println(car, year)

    // use this type of multiple declaration for better readability
	var (
		salary float64
		firstName string
		gender bool
	)

	fmt.Println(salary, firstName, gender)

	var a, b, c int
	fmt.Println(a, b, c)

	var i, j int
	i, j = 5, 10 // multiple assignment

	j, i = i, j // swap variables with multiple assignments
	fmt.Println(i, j)

	sum := 2 + 6 // expression in short declaration
	fmt.Println(sum)
}