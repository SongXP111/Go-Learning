package main

import "fmt"

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
func (p Phone) start() {
	fmt.Println(p.Name, "启动")
}

func (p Phone) stop() {
	fmt.Println(p.Name, "关机")
}

// 相机实现USB接口
type Camera struct {
}

func (c Camera) start() {
	fmt.Println("相机启动")
}

func (c Camera) stop() {
	fmt.Println("相机关机")
}

func (c Camera) run() {
	fmt.Println("相机在运行")
}

// 电脑
type Computer struct {
}

// 传入Phone或者Camera的实例
// 相当于 var usb Usber = phone / camera
func (c Computer) work(usb Usber) {
	usb.start()
	usb.stop()
}

func interfaceDemo1() {
	// p := Phone{
	// 	Name: "华为手机",
	// }

	// var p1 Usber // golang中接口就是一个数据类型
	// p1 = p       // 表示手机实现USB接口
	// p1.start()
	// p1.stop()

	// c := Camera{}
	// var c1 Usber = c // 表示相机实现了USB接口
	// c1.start()
	// c1.start()
	// // c1不能调用run方法，只能用c
	// c.run()

	var computer = Computer{}
	var phone = Phone{
		Name: "小米",
	}
	var camera = Camera{}

	// 手机连电脑
	computer.work(phone)
	computer.work(camera)
}
