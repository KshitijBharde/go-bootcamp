package main

import "fmt"

func main() {
	s1 := "Hi there Go!"
	fmt.Println(s1)

	fmt.Println("He says: \"Hello\"")

	fmt.Println(`He says: "Hello"`)

	s2 := `I like \n Go!`
	fmt.Println(s2)

	fmt.Println("Price: 100\nBrand: Nike")
	fmt.Println(`
Price: 100
Brand: Nike
	`)

	fmt.Println(`C:\Users\Kshitij`)
	fmt.Println("C:\\users\\Kshitij")

	s3 := "I love " + "Go " + "programming"
	fmt.Println(s3 + "!")

	// getting an element (byte) of a string:
    fmt.Println("Element at index zero:", s3[0]) // => 73 (ascii code for I)
    //  a string is in fact a slice of bytes in Go
 
    // strings are immutable and can't be changed
    // s3[5] = 'x' // => error: Cannon assign to s3[5]
}