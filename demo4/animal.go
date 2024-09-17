package main

import "fmt"

type Animaler1 interface {
	SetName(string)
}

type Animaler2 interface {
	GetName() string
}

// 接口嵌套
type Animaler3 interface {
	Animaler1
	Animaler2
}

type Dog struct {
	Name string
}

func (d *Dog) SetName(name string) {
	d.Name = name
}

func (d Dog) GetName() string {
	return d.Name
}

// type Cat struct {
// 	Name string
// }

// func (c *Cat) SetName(name string) {
// 	c.Name = name
// }

// func (c Cat) GetName() string {
// 	return c.Name
// }

func animalDemo() {
	var d1 = &Dog{
		Name: "aHei",
	}
	// var d2 Animaler = d1
	// d2.SetName("aHuang")
	// fmt.Println(d2.GetName())

	// var c1 = &Cat{}
	// var c2 Animaler = c1
	// c2.SetName("aTang")
	// fmt.Println(c2.GetName())

	// Dog实现两个接口
	var d2 Animaler1 = d1
	var d3 Animaler2 = d1

	d2.SetName("aHuang")
	fmt.Println(d3.GetName())

	// Dog实现接口嵌套
	var d4 Animaler3 = d1
	d4.SetName("aQi")
	fmt.Println(d4.GetName())
}
