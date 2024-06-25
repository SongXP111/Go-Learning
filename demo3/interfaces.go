package main

import "fmt"

// 接口是一个规范
type Usber interface {
	start()
	stop()
}

type Phone struct {
	Name string
}

func (p Phone) start() {
	fmt.Println(p.Name, "启动")
}

func (p Phone) stop() {
	fmt.Println(p.Name, "关机")
}

type Camera struct {
}

func (c Camera) start() {
	fmt.Println("相机 启动")
}

func (c Camera) stop() {
	fmt.Println("相机 关机")
}

// 照相机还可以有自己的方法
func (c Camera) run() {
	fmt.Println("相机 运行")
}

func interfaceDemo1() {
	p := Phone{
		Name: "华为",
	}
	// p.start()

	var p1 Usber // golang中接口是一个数据类型
	p1 = p       // 表示手机实现usb接口
	p1.start()
	p1.stop()

	c := Camera{}
	var c1 Usber = c // 相机实现了usb接口
	c1.start()
	c1.stop()
	c.run() // 接口不能调用相机的run方法
}
