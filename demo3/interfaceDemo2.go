package main

import "fmt"

/*
	空接口，表示没有任何约束

任意类型都可以实现空接口
空接口也可以当做类型来使用，可以表示任意类型
*/
type A interface {
}

// 可以当参数传入方法
func show(a interface{}) {
	fmt.Printf("value: %v, type: %T\n", a, a)
}

func interfaceDemo2() {
	/* var a A
	var str = "Hello World"
	a = str // 让字符串实现A这个接口
	fmt.Printf("value: %v, type: %T\n", a, a)

	var num = 20
	a = num // 让int类型实现A这个接口
	fmt.Printf("value: %v, type: %T\n", a, a)

	var flag = true
	a = flag // 让bool类型实现这个接口
	fmt.Printf("value: %v, type: %T\n", a, a) */

	/* var a interface{}
	a = 20
	fmt.Printf("value: %v, type: %T\n", a, a)
	a = "Hello World"
	fmt.Printf("value: %v, type: %T\n", a, a)
	a = true
	fmt.Printf("value: %v, type: %T\n", a, a) */

	show(20)
	show("Hello Golang")
	show(false)
	slice := []int{1, 2, 3, 4}
	show(slice)
}
