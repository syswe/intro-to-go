package main

import "fmt"

func main() {
	// Go Operators
	// + - * / % ++ --

	a := 10
	b := 20
	c := a + b
	d := a - b
	e := a * b
	f := a / b
	g := a % b
	a++
	b--
	fmt.Println(c, d, e, f, g, a, b)

	fmt.Println("-------------------------------")

	// == != < > <= >=
	a = 10
	b = 20
	c = 10
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a < b)
	fmt.Println(a > b)
	fmt.Println(a <= b)
	fmt.Println(a >= b)
	fmt.Println(a == c)

	fmt.Println("-------------------------------")

	// && || !
	a = 10
	b = 20
	c = 10
	fmt.Println(a == b && a == c)
	fmt.Println(a == b || a == c)
	fmt.Println(!(a == b))

	fmt.Println("-------------------------------")

	// & | ^ << >>

	a = 10 // 1010
	b = 20 // 10100
	c = 10 // 1010
	fmt.Println(a & b)
	// 1010
	fmt.Println(a | b)
	// 1010
	fmt.Println(a ^ b)
	// 1010
	fmt.Println(a << 1)
	// 10100
	fmt.Println(a >> 1)
	// 101

	fmt.Println("-------------------------------")

	// += -= *= /= %= &= |= ^= <<= >>=
	a = 10
	b = 20
	c = 10
	a += b
	fmt.Println(a)
	a -= b
	fmt.Println(a)
	a *= b
	fmt.Println(a)
	a /= b
	fmt.Println(a)
	a %= b
	fmt.Println(a)
	a &= b
	fmt.Println(a)
	a |= b
	// a = a | b
	fmt.Println(a)
	a ^= b
	//  a = a ^ b
	fmt.Println(a)
	a <<= 1
	// a = a << b
	fmt.Println(a)
	a >>= 1
	//  a = a >> b
	fmt.Println(a)

	// --------------------------------------------
}
