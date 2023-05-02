package main

import "fmt"

func main() {
	var employees map[string]string
	fmt.Printf("%#v\n", employees)

	fmt.Printf("No of pairs: %d\n", len(employees))

	fmt.Printf("The value for key %q is %q\n", "Kshitij", employees["Kshitij"])

	var accounts map[string]float64
	fmt.Printf("%#v\n", accounts["0x323"])

	// employees["Kshitij"] = "Programmer" // throws error, cannot assign key vals to unintialized maps

	people := map[string]float64{}

	people["Kshitij"] = 21.4
	people["Aditya"] = 10
	fmt.Println(people)

	map1 := make(map[int]int)
	map1[4] = 8

	balances := map[string]float64{
		"USD": 323.11,
		"EUR": 432.13,
		"CHF": 3234.1,
	}
	fmt.Println(balances)

	balances["USD"] = 500.2
	balances["GBP"] = 100
	fmt.Println(balances)

	fmt.Println(balances["RON"])

	balances["RON"] = 0

	v, ok := balances["RON"]

	if ok {
		fmt.Println("The RON value is: ", v)
	} else {
		fmt.Println("The RON key does not exist on map")
	}


	for k, v := range balances {
		fmt.Printf("Key: %#v, Value: %#v\n", k, v)
	}

	delete(balances, "USD")
	fmt.Println(balances)
}