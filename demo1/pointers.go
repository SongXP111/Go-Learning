package main

import "fmt"

func pointerDemo1() {
	// a对应了一个内存地址
	// a := 10
	// fmt.Println("a:")
	// fmt.Printf("Value: %v\n", a)
	// fmt.Printf("Type: %T\n", a)
	// fmt.Printf("Address: %p\n", &a)

	// 指针也是变量，是一种特殊的变量，储存着另一个对象的地址
	// p := &a
	// fmt.Println("p:")
	// fmt.Printf("Value: %v\n", p)
	// fmt.Printf("Type: %T\n", p)
	// fmt.Printf("Value: %p\n", &p)

	// 取指针的值
	a := 10
	p := &a // *int

	fmt.Println(p)  // a的地址
	fmt.Println(*p) // 取出p对应的内存地址的值

	*p = 20
	fmt.Println(p)
	fmt.Println(*p) // *p可以改变p对应地址的值

	d := 1
	pointerDemo2(&d)
	fmt.Println(d)

}

// 可以改变值
func pointerDemo2(a *int) {
	*a += 1
}
