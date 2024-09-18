package main

import (
	"fmt"
	"unsafe"
)

func intDemo() {
	var num int = 10
	fmt.Printf("num = %v, type:%T\n", num, num) // v: 原样输出
	fmt.Println(unsafe.Sizeof(num))             // 8 bits
}

func floatDemo() {
	var a float32 = 3.12
	fmt.Printf("a: %.2f, type: %T\n", a, a)
	fmt.Println(unsafe.Sizeof(a))

	// float精度丢失
	// 解决方法：第三方包

}

func boolDemo() {
	var b bool
	fmt.Println(b)
}

func stringDemo() {
	s1 := "Hello"
	fmt.Println(s1)

	// 多行字符串
	// str1 := `a
	// a
	// a
	// `

}

func charDemo() {
	// byte: ASCII
	// rune: utf8 (处理汉字)
	var a = 'a'
	fmt.Printf("Value: %v, type: %T\n", a, a)
	fmt.Printf("Value: %c, type: %T\n", a, a)

	var str = "this"
	fmt.Printf("Value: %v, type: %T\n", str[2], str[2]) // i

	s := "你好 golang"
	// rune
	// 如果用for i 循环，输出的是byte
	for _, v := range s {
		fmt.Printf("%v(%c)", v, v)
	}

	// 修改字符串
	// 先转成[]byte(s)，修改后再转成string
	// 如果有汉字，需要用[]rune

}
