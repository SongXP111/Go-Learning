package main

import "fmt"

// defer逆序执行

func deferDemo1() {
	fmt.Println("Start")
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("End")
}

func deferDemo2() {
	fmt.Println("Start")
	defer func() {
		fmt.Println("AAAA")
	}()
	fmt.Println("End")
}

// 返回0
// 匿名返回值
func deferDemo3() int {
	var a int
	defer func() {
		a++
	}()
	return a
}

// 返回1
// 命名返回值
func deferDemo4() (a int) {
	defer func() {
		a++
	}()
	return a
}
