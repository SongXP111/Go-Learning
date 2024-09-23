package main

import (
	"fmt"
	"reflect"
)

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	kind := v.Kind()
	switch kind {
	case reflect.Int:
		fmt.Printf("int类型的原始值%v\n", v.Int())
	case reflect.Float32:
		fmt.Printf("float32类型的原始值%v\n", v.Float())
	case reflect.Float64:
		fmt.Printf("float64类型的原始值%v\n", v.Float())
	case reflect.String:
		fmt.Printf("string类型的原始值%v\n", v.String())
	default:
		fmt.Printf("还没有判断这个类型\n")
	}
}

func main() {
	var a float32 = 3.14
	reflectValue(a)

	var b int64 = 100
	reflectValue(b)

	var c string = "Hello Golang"
	reflectValue(c)
}
