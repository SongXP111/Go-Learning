package main

import (
	"fmt"
	"reflect"
)

// 通过反射修改变量的值
func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(123)
	} else if v.Elem().Kind() == reflect.String {
		v.Elem().SetString("Hello World")
	}
}

func main() {
	var a int64 = 100
	reflectSetValue(&a)
	fmt.Println(a)

	var b string = "Hello Golang"
	reflectSetValue(&b)
	fmt.Println(b)
}
