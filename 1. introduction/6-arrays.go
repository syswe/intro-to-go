package main

import "fmt"

// arrayler
func main() {
	// array_name:= [length]Type{index1: item1, index2: item2, index3: item3,...indexN: itemN}
	// array examples
	var a [5]int
	a[2] = 7
	a[4] = 12
	a = [5]int{0, 0, 7, 0, 12}
	fmt.Println(a)

	b := [5]int{5, 4, 3, 2, 1}
	b = [5]int{5, 4, 3, 2, 1}
	fmt.Println(b)

	c := [5]int{1, 2, 3}
	c = [5]int{1, 2, 3, 0, 0}
	fmt.Println(c)

	d := [...]int{4, 5, 6, 7, 8, 9}
	d = [6]int{4, 5, 6, 7, 8, 9}
	fmt.Println(d)

	f := [...]int{1: 10, 2: 20, 4: 50}
	fmt.Println(f)
	f = [5]int{0, 10, 20}
	fmt.Println(f, len(f), cap(f), f[0], f[1], f[2])

	var twoD [2][3]int
	twoD = [2][3]int{{0, 0, 0}, {0, 0, 0}}
	fmt.Println(twoD)

	fmt.Println("-----------------------")
	// var threeD [2][3][4]int
	threeD := [2][3][4]int{{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}, {{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}}
	fmt.Println(threeD, len(threeD), len(threeD[0]), len(threeD[0][0]))

	// arraylerin kopyalanması
	var arr1 [5]int
	arr2 := arr1
	arr2[0] = 10
	fmt.Println(arr1, arr2)

	fmt.Println("-----------------------")
	// arraylerin karşılaştırılması
	arr3 := [3]int{1, 2, 3}
	arr4 := [3]int{1, 2, 3}
	arr5 := [3]int{1, 2, 4}
	fmt.Println(arr3 == arr4, arr3 == arr5, arr4 == arr5)

	fmt.Println("-----------------------")
	// arraylerin pointerları
	arr8 := [3]int{1, 2, 3}
	arr9 := &arr8
	fmt.Println(arr8, arr9)
	arr8[0] = 10
	fmt.Println(arr8, arr9)
	arr9[1] = 20
	fmt.Println(arr8, arr9)

	fmt.Println("-----------------------")
	// arraylerin slice'ları
	arr10 := [3]int{1, 2, 3}
	slice1 := arr10[:]
	slice2 := arr10[1:]
	slice3 := arr10[:2]
	slice4 := arr10[1:2]
	fmt.Println(arr10, slice1, slice2, slice3, slice4)
	arr10[1] = 20
	fmt.Println(arr10, slice1, slice2, slice3, slice4)
	slice1[2] = 30
	fmt.Println(arr10, slice1, slice2, slice3, slice4)
	slice2[0] = 40
	fmt.Println(arr10, slice1, slice2, slice3, slice4)
	slice3[1] = 50
	fmt.Println(arr10, slice1, slice2, slice3, slice4)
	slice4[0] = 60
	fmt.Println(arr10, slice1, slice2, slice3, slice4)

}
