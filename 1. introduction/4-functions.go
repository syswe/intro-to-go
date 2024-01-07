package main

import "fmt"

func main() {
	// 2. functions
	result := add(10, 20)
	result2 := cikar(20, 10)
	fmt.Println(result)
	fmt.Println(result2)

	// --------------------------------------------
}

func add(a int, b int) int {
	return a + b
}

func (a int, b int) int {
	return a - b
}
