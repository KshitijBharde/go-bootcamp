package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := string(99)
	fmt.Println(s) // ascii representation of 99

	str := fmt.Sprintf("%f", 44.2)
	fmt.Println(str)

	str1 := fmt.Sprintf("%d", 32425)
	fmt.Println(str1)
	fmt.Println(string(32425)) // => unicode character for number 32425

	sf1 := "3.123"
	fmt.Printf("%T\n", sf1)

	var f1, err = strconv.ParseFloat(sf1, 64)
	_ = err
	fmt.Printf("%T\n", f1)

	i, err := strconv.Atoi("-45")
	s2 := strconv.Itoa(34)

	fmt.Printf("i type is %T, i value is %v\n", i, i)
	fmt.Printf("s2 type is %T, s2 value is %q\n", s2, s2)

}