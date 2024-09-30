package main

import "fmt"

// make和new

func makeDemo1() {
	// 引用数据类型必须使用make分配内存
	// make用于slice，map，和channel

	// var userinfo = make(map[string]string)
	// userinfo["username"] = "Xipeng"
	// fmt.Println(userinfo)

	// var a = make([]int, 4)
	// a[0] = gin1
	// fmt.Println(a)

	var a = new(int) // pointer也是引用数据类型
	*a = 100
	fmt.Println(*a)

}

func newDemo1() {
	// new返回的是指针类型
	var a = new(int) // a是一个指针变量，类型是*int的指针类型，值是0
	fmt.Printf("Value: %v\n", a)
	fmt.Printf("Type: %T\n", a)
	fmt.Printf("Value at this address: %v\n", *a)
}
