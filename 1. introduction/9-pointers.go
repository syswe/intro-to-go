package main

// pointers
func main() {
	// pointers
	var a int = 42
	var b *int = &a
	println(a, b)
	println(a, *b)
	a = 27
	println(a, *b)
	*b = 14
	println(a, *b)

	// new
	var c *int = new(int)
	println(c)
	*c = 10
	println(*c)

	// make
	var d *[]int = new([]int)
	println(d)
	*d = make([]int, 10, 100)
	println(d)
	println(len(*d), cap(*d))

	// pointers
	var e []int = make([]int, 10, 100)
	println(e)
	println(len(e), cap(e))

	var f *[]int = &e
	println(f)
	println(len(*f), cap(*f))

	// pointers
	var g map[string]int = make(map[string]int)
	println(g)

	var h *map[string]int = &g
	println(h)

}
