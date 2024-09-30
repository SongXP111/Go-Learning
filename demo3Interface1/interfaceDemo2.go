package main

import "fmt"

/*
	空接口，表示没有任何约束

任意类型都可以实现空接口
空接口也可以当做类型来使用，可以表示任意类型
*/
type A interface {
}

type B struct {
	Name string
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

	// 可以随意赋值
	/* var a interface{}
	a = 20
	fmt.Printf("value: %v, type: %T\n", a, a)
	a = "Hello World"
	fmt.Printf("value: %v, type: %T\n", a, a)
	a = true
	fmt.Printf("value: %v, type: %T\n", a, a) */

	// 作为方法的参数
	/* show(20)
	show("Hello Golang")
	show(false)
	slice := []int{gin1, 2, 3, 4}
	show(slice) */

	// 可以作为map的value的类型
	/* var m1 = make(map[string]interface{})
	m1["username"] = "张三"
	m1["age"] = 20
	m1["married"] = false
	fmt.Println(m1) */

	// 也可以放在切片里
	/* var s1 = []interface{}{20, "hello Golang", true}
	fmt.Println(s1) */

	// 类型断言
	// 经常结合switch语句使用
	var a interface{}
	a = "hello Golang"
	v, ok := a.(string)
	if ok {
		fmt.Println("a就是string类型，值是: ", v)
	} else {
		fmt.Println("断言失败")
	}
}

/*
定义一个方法
可以传入然后类型
根据不同的类型实现不同的功能
*/
func MyPrint1(x interface{}) {
	if _, ok := x.(string); ok {
		fmt.Println("String")
	} else if _, ok := x.(int); ok {
		fmt.Println("Integer")
	} else if _, ok := x.(bool); ok {
		fmt.Println("Boolean")
	} else {
		fmt.Println("Unknown Type")
	}
}

// x.(type)判断一个变量的类型，只能用在switch语句里面
func MyPrint2(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("Integer")
	case string:
		fmt.Println("String")
	case bool:
		fmt.Println("Boolean")
	case B:
		fmt.Println("B struct") // 也可以判断结构体
	default:
		fmt.Println("Unsupported Type")
	}
}
