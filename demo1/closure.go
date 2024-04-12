package main

import "fmt"

// 闭包
// 函数里面嵌套一个函数，然后返回里面的函数
// 可以让一个变量常驻内存，且不污染全局
func adder1() func() int {
	var i = 10
	return func() int {
		return i + 1
	}
}

func adder2() func(y int) int {
	var i = 10
	return func(y int) int {
		i += y
		return i
	}
}

func adderTest() {
	var fn1 = adder1() //执行
	fmt.Println(fn1()) // 11
	fmt.Println(fn1()) // 11
	fmt.Println(fn1()) // 11

	var fn2 = adder2()   //执行
	fmt.Println(fn2(10)) // 20
	fmt.Println(fn2(10)) // 30
	fmt.Println(fn2(10)) // 40
}
