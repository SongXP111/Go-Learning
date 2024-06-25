package main

import (
	"fmt"
	"modDemo/calc"
	T "modDemo/tools" // T是包别名
)

func main() {
	sum := calc.Add(10, 2)
	fmt.Println(sum)
	fmt.Println(calc.Age)

	b := T.Mul(4, 5)
	fmt.Println(b)
}

// 在main包中，init在main之前执行
func init() {
	fmt.Println("main init...")
}
