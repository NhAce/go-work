package main

import (
	"fmt"
)
//fibonacci: 1,1,2,3,5,8,13
func fibonacci() func(int) int {
	a := 1
	b := 0
	return func(x int) int {
		if x > 0 {
			c := a
			a += b
			b = c
		}
		return a
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}

