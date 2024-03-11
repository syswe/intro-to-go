package main

import "fmt"

func main() {
	// Creating a map
	studentGrades := make(map[string]int)

	// Adding key-value pairs to the map
	studentGrades["Alice"] = 90
	studentGrades["Bob"] = 85
	studentGrades["Charlie"] = 95

	// Accessing values from the map
	aliceGrade := studentGrades["Alice"]
	fmt.Println("Alice's grade:", aliceGrade)

	// Updating values in the map
	studentGrades["Bob"] = 90

	// Deleting a key-value pair from the map
	delete(studentGrades, "Charlie")

	// Iterating over the map
	for name, grade := range studentGrades {
		fmt.Println(name, ":", grade)
	}
}
