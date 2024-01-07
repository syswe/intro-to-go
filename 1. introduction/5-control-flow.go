package main
// 4. kontrol akışı örnekleri
import "fmt"

main() {
	
	// if, else, else if
	a := 10
	b := 20
	if a == b {
		fmt.Println("a == b")
	} else if a < b {
		fmt.Println("a < b")
	} else {
		fmt.Println("a > b")
	}

	// for, range
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for _, deger := range []int{1, 2, 3, 4, 5} {
		fmt.Println(deger)
	}

	// while
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// switch, case, default
	switch foo := 2; foo {
	case 1:
		fmt.Println("foo 1")
	case 2:
		fmt.Println("foo 2")
	default:
		fmt.Println("foo 1 ve 2 değil")
	}

	fmt.Println("--------------------------")
	// break, continue, goto
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	
	fmt.Println("--------------------------")
	// defer
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")
	fmt.Println("main function")



}