package main

import "fmt"

// Define a struct type
type Rectangle struct {
	width, height float64
}

// Define a method for the Rectangle type
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Define another method for the Rectangle type
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func main() {
	// Create a new Rectangle object
	rect := Rectangle{width: 10, height: 5}

	// Call the methods on the Rectangle object
	fmt.Println("Area:", rect.Area())
	fmt.Println("Perimeter:", rect.Perimeter())
}
