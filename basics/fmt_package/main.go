package main

import "fmt"

func main() {
	name := "Kshitij"
	fmt.Println("Hello, I am", name)

	a, b := 4, 6
	fmt.Println("Sum:", a+b, "Mean Value:", (a+b)/2)

	fmt.Printf("Your age is %d\n", 28)
	fmt.Printf("He says \"Hello Go!\"\n")

	figure := "Circle"
	radius := 5
	pi := 3.14159
	diameter :=  float64(radius)*2*pi

	_, _ = figure, pi

	fmt.Printf("Radius is %d\n", radius)
	fmt.Printf("Radius is %+d\n", radius) // add + to show sign

	fmt.Printf("Pi constant is %f\n", pi)

	fmt.Printf("The diameter of a %s with a Radius of %d is %f\n", figure, radius, diameter)

	//%q for quoted string
	fmt.Printf("This is a %q\n", figure)

	//%v -> replaced by any type
	fmt.Printf("The diameter of a %v with a Radius of %v is %v\n", figure, radius, diameter)

	//%T -> type of variable
	fmt.Printf("figure is of type %T\n", figure)
	fmt.Printf("radius is of type %T\n", radius)
	fmt.Printf("pi is of type %T\n", pi)

	//%t -> boolean
	closed := true
	fmt.Printf("File closed: %t\n", closed)

	//%b -> binary/base 2
	x := 55
	fmt.Printf("%d in binary: %b\n", x, x)
    fmt.Printf("%d in binary(8bits/byte): %08b\n", x, x)

	//%x -> hexadecimal
	y := 100
	fmt.Printf("%d in hexadecimal: %x\n", y, y)

	z := 3.4
	k := 6.9
    
	// Limiting decimal points
	fmt.Printf("z * k = %.2f\n", z*k)
}