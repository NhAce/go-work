package main

import (
	"fmt"
)

func main()  {
	s := []int{2, 3, 5, 7, 9, 11, 13}
	fmt.Println("s == ", s)
	fmt.Println("s[1:4] == ", s[1:4])
	//省略下表代表从 0 开始
	fmt.Println("s[:3] == ", s[:3])
	//省略上标代表到 len(s) 结束
	fmt.Println("s[2:] == ", s[2:])
}

