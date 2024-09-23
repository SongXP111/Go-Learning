package main

import (
	"fmt"
	"reflect"
)

type myInt int
type Person struct {
	Name string
	Age  int
}

// 反射获取任意变量的类型
func reflectFn(x interface{}) {
	v := reflect.TypeOf(x)
	//v.Name() // 类型名称
	//v.Kind() // 种类（底层类型）
	fmt.Printf("类型: %v, 类型名称: %v, 类型种类:%v\n", v, v.Name(), v.Kind())
}

func main() {
	a := 10
	b := 23.4
	c := true
	d := "Hello Golang"

	reflectFn(a) // int
	reflectFn(b) // float64
	reflectFn(c) // bool
	reflectFn(d) // string

	var e myInt = 34
	var f = Person{
		Name: "zhangsan",
		Age:  20,
	}
	reflectFn(e) // main.myInt
	reflectFn(f) // main.Person

	var h = 25
	reflectFn(&h) // *int

	var i = [3]int{1, 2, 3}
	reflectFn(i) // array

	var j = []int{1, 2, 3}
	reflectFn(j) // slice
}
