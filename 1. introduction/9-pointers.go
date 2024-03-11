package main

import "fmt"

func main() {
	// Declare a variable
	var num int = 10

	// Declare a pointer variable
	var ptr *int

	// Assign the address of num to the pointer variable
	ptr = &num

	// Print the value and address of num
	fmt.Println("Value of num:", num)
	fmt.Println("Address of num:", &num)

	// Print the value and address stored in the pointer variable
	fmt.Println("Value stored in ptr:", *ptr)
	fmt.Println("Address stored in ptr:", ptr)

	// Change the value of num using the pointer variable
	*ptr = 20

	// Print the updated value of num
	fmt.Println("Updated value of num:", num)
}
