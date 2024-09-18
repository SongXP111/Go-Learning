package main

import (
	"fmt"
	"time"
)

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("Test: Hello Golang")
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go test() // 表示开启一个协程，并行执行
	for i := 0; i < 10; i++ {
		fmt.Println("Main: Hello Golang")
		time.Sleep(time.Millisecond * 100)
	}
}
