package main

import "fmt"

// Define a struct type
type Person struct {
	Name    string
	Age     int
	Country string
}

func main() {
	// Create a new instance of the Person struct
	person := Person{
		Name:    "John Doe",
		Age:     30,
		Country: "USA",
	}

	// Access and modify struct fields
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
	fmt.Println("Country:", person.Country)

	// Update struct fields
	person.Age = 31
	person.Country = "Canada"

	fmt.Println("Updated Age:", person.Age)
	fmt.Println("Updated Country:", person.Country)
}
