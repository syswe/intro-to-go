package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{Name: "John", Age: 30}

	// Get the value of the "Name" field using reflection
	nameField := reflect.ValueOf(person).FieldByName("Name")
	fmt.Println("Value of Name field:", nameField)

	// Set the value of the "Age" field using reflection
	ageField := reflect.ValueOf(&person).Elem().FieldByName("Age")
	ageField.SetInt(40)
	fmt.Println("Updated value of Age field:", person.Age)
}
