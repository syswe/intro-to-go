// Examples of variables and data types in go language
package main

import "fmt"

func main() {

	// var i int = 10
	// var f float32 = 3.14
	// var s string = "Hello"
	// var b bool = true

	i := 10
	f := 3.14
	s := "Hello"
	b := true

	var t int // z is initialized to 0

	var c, d, _ int = 10, 20, 30 // _ is a blank identifier
	var e, fl, g = 10, 3.14, "Hello"
	var (
		x = 10
		y = 20
		z = 30
	)
	var (
		xx int     = 10
		yy float32 = 3.14
		zz string  = "Hello"
	)

	fmt.Println(i, f, s, b, t)
	fmt.Println(c, d)
	fmt.Println(e, fl, g)
	fmt.Println(x, y, z)
	fmt.Println(xx, yy, zz)

	// int8, int16, int32, int64
	// uint8, uint16, uint32, uint64
	// float32, float64
	// complex64, complex128
	// byte, rune
	// string
	// bool
	fmt.Println("-------------------------------")

	entry := 74
	// %d is for decimal
	// %b is for binary
	// %o is for octal
	// %x is for hexadecimal
	// %v is for any value
	// %T is for type
	// %t is for boolean
	// %s is for string
	// %q is for quoted string
	// %f is for float
	// %e is for scientific notation
	// %E is for scientific notation
	// %c is for character
	// %p is for pointer
	fmt.Printf("%d ==  %v\n", entry, entry)
	fmt.Printf("%d type is %T\n", entry, entry)
	fmt.Printf("%d in binary is %b\n", entry, entry)
	fmt.Printf("%d in octal is %o\n", entry, entry)
	fmt.Printf("%d in hexadecimal is %x\n", entry, entry)
}
