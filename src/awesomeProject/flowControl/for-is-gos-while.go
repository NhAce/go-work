package main

import (
	"fmt"
)

func main()  {
	sum := 1
	// for 在此处的功能和 while 一样
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
