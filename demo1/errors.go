package main

import (
	"errors"
	"fmt"
)

// panic：直接结束
// recover：只能用在defer函数

func panic1() {
	fmt.Println("1")
}

func panic2() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	panic("Error")
}

func panic3(a, b int) int {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Error: ", err) // Error:  runtime error: integer divide by zero
		}
	}()
	return a / b
}

func readFile(fileName string) error {
	if fileName == "main.go" {
		return nil
	} else {
		return errors.New("读取文件失败")
	}
}

func readFileTest() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("给管理员发送邮件")
		}
	}()
	err := readFile("xxx.go")
	if err != nil {
		panic(err)
	}
}

func panicDemo() {
	// panic1()
	// panic2()
	// fmt.Println("End")

	// fmt.Println(panic3(10, 0))
	// fmt.Println("End")
	// fmt.Println(panic3(10, 2))
	readFileTest()
	fmt.Println("End")
}
