package main

import (
	"fmt"
	"math"
)


func main()  {
	z := 1.0
	var m float64 = 0.0
	for x := 1; x <= 100; x++ {
		value := z - ((math.Pow(z, 2) - float64(x)) / float64((2 * x)))
		fmt.Printf("x is %v, value is %v\n", x, value)
		if math.Abs(value - m) < 0.0001{
			fmt.Printf("x is %v, value is %v, m is %v\n", x, value, m)
			break
		}else{
			m = value
		}

		//fmt.Printf("math.Sqrt(%v) is %v\n", x, math.Sqrt(float64(x)))

	}
}