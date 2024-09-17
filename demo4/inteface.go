package main

import "fmt"

// 结构体值接收者实现接口：
// 值接收者：如果结构体中的方法是值接收者，那么实例化后的结构体值类型和结构体指针类型都可以赋值给接口变量

// 接口相当于一个规范
// golang的接口是一种数据类型，不需要显式实现
type Usber interface {
	start()
	stop()
}

// 如果接口里面有方法的话，必须要通过结构体或者自定义类型来实现
type Phone struct {
	Name string
}

// 手机要实现USB接口的话，必须实现所有方法
// func (p Phone) start() { // 值接收者
// 	fmt.Println(p.Name, "启动")
// }

func (p *Phone) start() { // 指针接收者
	fmt.Println(p.Name, "启动")
}

func (p *Phone) stop() {
	fmt.Println(p.Name, "关机")
}

func interfaceDemo() {
	// 结构体值接收者例化后的结构体值类型和结构体指针类型都可以复制给接口变量

	// var p1 = Phone{
	// 	Name: "小米手机",
	// }

	// var p2 Usber = p1 // 表示让Phone实现usb接口
	// p2.start()

	// var p3 = &Phone{
	// 	Name: "苹果手机",
	// }
	// var p4 Usber = p3
	// p4.start()

	// 传入指针
	var p1 = &Phone{ // 必须传入指针地址
		Name: "小米手机",
	}
	var p2 Usber = p1 // 表示让Phone实现usb接口
	p2.start()
}
