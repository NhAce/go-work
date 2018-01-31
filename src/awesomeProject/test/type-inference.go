package main

import (
	"fmt"
	"math"
)

func main()  {
	var i = math.Pi
	fmt.Printf("i is of type %T\n", i)
	v := 42
	fmt.Printf("v is of type %T\n", v)
	f := 63.1
	fmt.Printf("f is of type %T\n", f)
	b := false
	fmt.Printf("b is of type %T\n", b)
	s := "hello"
	fmt.Printf("s is of type %T\n", s)
}

