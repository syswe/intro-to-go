package main

// constants
const (
	PI       = 3.14
	Language = "Go"
)

// iota
const (
	Zero = iota
	One
	Two   // 2
	Three // 3
	Four  // 4
)

// iota
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB // 2^20
	GB // 2^30
	TB // 2^40
	PB // 2^50
	EB // 2^60
	ZB // 2^70
	YB // 2^80
)

// iota
const (
	_   = iota + 5
	KB2 = 1 << (10 * iota)
	MB2 // 2^20
	GB2 // 2^30
	TB2 // 2^40
)

func main() {

	// constants
	println(PI, Language)

	// iota
	println(Zero, One, Two, Three, Four)

	// iota
	println(KB, MB, GB, TB, PB, EB, ZB, YB)

	// iota
	println(KB2, MB2, GB2, TB2)

	// try change constants like: KB2 = 10
	// you see error: cannot assign to KB2

}
