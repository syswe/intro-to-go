package main

import "fmt"

// Define an interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Define a struct for a rectangle
type Rectangle struct {
	width  float64
	height float64
}

// Implement the Shape interface for Rectangle
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Define a struct for a circle
type Circle struct {
	radius float64
}

// Implement the Shape interface for Circle
func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}

func main() {
	// Create a rectangle
	rectangle := Rectangle{width: 5, height: 3}

	// Create a circle
	circle := Circle{radius: 2}

	// Use the Shape interface to calculate area and perimeter
	shapes := []Shape{rectangle, circle}
	for _, shape := range shapes {
		fmt.Println("Area:", shape.Area())
		fmt.Println("Perimeter:", shape.Perimeter())
		fmt.Println()
	}
}
