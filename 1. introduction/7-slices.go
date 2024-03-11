package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
func main() {

	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:6]
	fmt.Println(s)

	s = s[0:3]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

	s = s[1:4]
	fmt.Println(s)

	// ---
	s = []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	// nil slice
	var a []int
	fmt.Println(a, len(a), cap(a))
	if a == nil {
		fmt.Println("nil!")
	}

	// Creating a slice with make

	a = make([]int, 5)
	printSlice(a)

	b := make([]int, 0, 5)
	printSlice(b)

	c := b[:2]
	printSlice(c)

	d := c[2:5]
	printSlice(d)

	// Appending to a slice
	var e []int
	printSlice(e)

	// append works on nil slices.
	e = append(e, 0)
	printSlice(e)

	// add more advanced slices examples deep-dive and detailed like real life examples add much more detailed example

	// Slices of slices
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"}}
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", board[i])
	}

	// Appending to a slice
	var f []int
	printSlice(f)

	// append works on nil slices.
	f = append(f, 0)
	printSlice(f)

	// The slice grows as needed.
	f = append(f, 1)
	printSlice(f)

	// We can add more than one element at a time.
	f = append(f, 2, 3, 4)
	printSlice(f)

	// Range
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// You can skip the index or value by assigning to _.
	pow = make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {

		fmt.Printf("%d\n", value)
	}

}
