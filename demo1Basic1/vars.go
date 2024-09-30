package main

import (
	"fmt"
)

var global = "global"

func varDemo1() {
	// 声明一个变量

	// var username string
	// fmt.Println(username) // 没有初始化的string值为空

	// var username string
	// username = "Xipeng"
	// fmt.Println(username)

	// var username string = "Xipeng"
	// fmt.Println(username)

	// var username = "Xipeng"
	// var age = 24
	// var sex = "Male"
	// fmt.Println(username, age, sex)

	// 一次声明多个变量
	// var a1, a2 string
	// a1 = "Xipeng"
	// a2 = "Mino"
	// fmt.Println(a1, a2)

	// 短变量声明法（只能声明局部变量，不能声明全局变量）
	username := "Xipeng"
	fmt.Println(username)

	a, b, c := 1, 2, 3
	fmt.Println(a, b, c)

	e, f, g := 4, 5, "6"
	fmt.Println(e, f, g)

	fmt.Println(global)

	// 常量：const
	const constant = 10
}

func varDemo2() {
	var username = "Xipeng"
	username = "Mino"

	fmt.Println((username))

	// 常量必须赋值
	const pi = 3.14159

	// 常量的值不可以改变
	// pi = 3.14 // 报错

	// 多个常量
	// 如果省略了，则表示和上面一行相同
	const (
		A = 100
		B = 10
		C // 10
		D = 20
	)

	fmt.Println(A, B, C, D)

	// iota
	// 自动累加
	const (
		n1 = iota // 0
		_         // gin1 跳过
		n3
		n4
	)

	fmt.Println(n1, n3, n4)
}
