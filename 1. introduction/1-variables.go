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

	var c, d int = 10, 20
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

	fmt.Println(i, f, s, b)
	fmt.Println(c, d)
	fmt.Println(e, fl, g)
	fmt.Println(x, y, z)
	fmt.Println(xx, yy, zz)

	fmt.Println("-------------------------------")

	// int8, int16, int32, int64
	// uint8, uint16, uint32, uint64
	// float32, float64
	// complex64, complex128
	// byte, rune
	// string
	// bool

}
