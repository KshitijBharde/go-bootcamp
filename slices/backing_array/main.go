package main

import (
	"fmt"
)

func main() {
	s1 := []int{10, 20, 30, 40, 50}
	fmt.Println(s1)
	s3, s4 := s1[0:2], s1[1:3]
	fmt.Println(s3, s4)

	s3[1] = 600

	fmt.Println(s1, s3, s4)

	arr1 := [4]int{10, 20, 30, 40}
	slice1, slice2 := arr1[:2], arr1[1:3]

	arr1[1] = 2
	fmt.Println(arr1, slice1, slice2)

	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	newCars = append(newCars, cars[:2]...)

	cars[0] = "Nissan"
	fmt.Println(cars, newCars)

	s5 := []int{10, 20, 30, 40, 50}
	newSlice := s5[0:3]
	fmt.Println(len(newSlice), cap(newSlice))

	newSlice = s5[2:5] // {30, 40, 50}
	fmt.Println(len(newSlice), cap(newSlice))

	fmt.Printf("%p\n", &s5)

	fmt.Printf("%p %p\n", &s5, &newSlice)

	newSlice[0] = 100
	fmt.Println("s5:", s5)
}