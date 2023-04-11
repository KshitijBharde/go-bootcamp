package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5}

	b := a[1:3]
	fmt.Printf("%v, %T\n", b, b) 

	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1[1:3]
	fmt.Println(s2)

	fmt.Println(s1[2:]) // same as s1[2:len(s1)]
	fmt.Println(s1[:3]) // same as s1[0:3]
	fmt.Println(s1[:]) // same as s1[0:len(s1)]
	fmt.Println(s1[:100]) // gives error: slice bounds out of range

	s1 = append(s1[:4], 100) // => [1, 2, 3, 4, 100]
	s1 = append(s1[:4], 200) // => [1, 2, 3, 4, 200]
}