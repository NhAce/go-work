package main

import (
	"fmt"
)

func main()  {
	sum := 1
	//for i:= 0; i < 10; i++ {
	//	sum += 1
	//}
	for ; sum < 1000;  {
		sum += sum
	}
	fmt.Println(sum)
}

