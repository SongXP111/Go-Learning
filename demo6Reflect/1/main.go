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
	fmt.Println(v)
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
}
