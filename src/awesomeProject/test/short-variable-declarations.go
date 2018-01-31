package main

import (
	"fmt"
)
// 函数外的每个语句必须以关键字开始（var, func, const, type, import 等），
//m := 3
func main()  {
	var i, j = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}

