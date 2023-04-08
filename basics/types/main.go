package main

import (
	"fmt"
)

func main() {
	// uint8      the set of all unsigned  8-bit integers (0 to 255)
    // uint16      the set of all unsigned 16-bit integers (0 to 65535)
    // uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
    // uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)
 
    // int8        the set of all signed  8-bit integers (-128 to 127)
    // int16       the set of all signed 16-bit integers (-32768 to 32767)
    // int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
    // int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
 
    // uint     either 32 or 64 bits
    // int      same size as uint
 
    // float32     the set of all IEEE-754 32-bit floating-point numbers
    // float64     the set of all IEEE-754 64-bit floating-point numbers
    // complex64   the set of all complex numbers with float32 real and imaginary parts
    // complex128  the set of all complex numbers with float64 real and imaginary parts

	// byte        alias for uint8
	// INT type
	var i1 int8 = -128
	fmt.Printf("%T\n", i1)

	var i2 uint16 = 65535
	fmt.Printf("%T\n", i2)

	// FLOAT type
	var f1, f2, f3 float64 = 1.1, -.2, 5.
	fmt.Printf("%T %T %T\n", f1, f2, f3)


	// RUNE type
	// rune        alias for int32
	var r rune = 'f'
	fmt.Printf("%T\n", r) // => int32 (rune is an alias to int32)
	fmt.Printf("%v\n", r) // decimal ascii code
	fmt.Printf("%x\n", r) // => 66,  the hexadecimal ascii code for 'f'
	fmt.Printf("%c\n", r) // => f

	// BOOL type
	var b bool = true
	fmt.Printf("%T\n", b)  // => bool

	// STRING type
	var s string = "Hello Go!"
	fmt.Printf("%T\n", s) // => string

	// ARRAY type
	var nums = [4]int{4, 5, 6, -100}
	fmt.Printf("%T\n", nums) // =>  [4]int

	// SLICE type
	var cities = []string{"London", "Tokyo", "New York"}
	fmt.Printf("%T\n", cities) // => []string

	// MAP type
	balances := map[string]float64{
		"USD": 2332.2,
		"EUR": 511.11,
	}
	fmt.Printf("%T\n", balances) // => map[string]float64

	// STRUCT type
	type Person struct{
		name string
		age int
	}
	var you Person
	fmt.Printf("%T\n", you) // => main.Person

	// POINTER type
	var x int = 2
	ptr := &x
	fmt.Printf("ptr is of type %T with value of %v\n", ptr, ptr) // => ptr is of type *int with value 0xc0000b6028(hexadecimal)

	// FUNCTION type
	fmt.Printf("%T\n", f)
}

func f() {}