package main

import (
	"fmt"
)

type User struct {
	name string
	age int
	gender string
}

func main()  {
	u := User{"qxl", 27, "male"}
	//u.age = 28
	//fmt.Println(u)
	p := &u
	fmt.Println(p)
	p.age = 18
	fmt.Println(*p)
	fmt.Println(u)
}
