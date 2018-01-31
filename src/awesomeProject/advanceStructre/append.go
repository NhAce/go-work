package main

import (
	"fmt"
)

func main() {
	var a []int
	printSlice1("a", a)

	a = append(a, 0)
	printSlice1("a", a)

	a = append(a, 1)
	printSlice1("a", a)

	a = append(a, 2, 3, 4, 5)
	printSlice1("a", a)

}

func printSlice1(s string, x []int)  {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
