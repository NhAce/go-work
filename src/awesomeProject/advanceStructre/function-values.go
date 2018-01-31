package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	//直接调用 hypot 方法，返回 math.Sqrt(5*5 + 12*12)
	fmt.Println(hypot(5,12))
	//调用 compute 方法，hypot 方法作为 compute 的参数传入，最终相当执行 hypot(3, 4)
	fmt.Println(compute(hypot))
	//调用 compute 方法，math.Pow 方法作为 compute 的参数传入，最终相当执行 math.Pow(3, 4)
	fmt.Println(compute(math.Pow))
}

