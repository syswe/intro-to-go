package main

import (
	"fmt"
	"time"
)

// goroutines
func main() {
	go printNumbers()
	go printLetters()

	// Wait for goroutines to finish
	time.Sleep(2 * time.Second)
}

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func printLetters() {
	for i := 'a'; i <= 'e'; i++ {
		fmt.Printf("%c ", i)
		time.Sleep(500 * time.Millisecond)
	}
}
