package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email,omitempty"`
}

func main() {
	// Marshaling JSON
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john.doe@example.com",
	}

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	fmt.Println(string(jsonData))

	// Unmarshaling JSON
	jsonData = []byte(`{"name":"Jane Smith","age":25}`)

	var anotherPerson Person
	err = json.Unmarshal(jsonData, &anotherPerson)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	fmt.Println(anotherPerson.Name)
	fmt.Println(anotherPerson.Age)
	fmt.Println(anotherPerson.Email)
}
