package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func test1() {
	for i := 0; i < 10; i++ {
		fmt.Println("Test1: Hello Golang - ", i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done() // 协程计数器-1
}

func test2() {
	for i := 0; i < 10; i++ {
		fmt.Println("Test2: Hello Golang - ", i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done() // 协程计数器-1
}

func main() {
	wg.Add(1)  // 协程计数器+1
	go test1() // 表示开启一个协程，并行执行
	wg.Add(1)
	go test2()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("Main: Hello Golang - ", i)
	// 	time.Sleep(time.Millisecond * 50) // 如果主进程执行的快，协程会提前结束
	// }
	// 解决方法：
	// 1. 休眠（不是好方法）
	// time.Sleep(time.Second)

	// 2. WaitGroup
	wg.Wait()
	fmt.Println("主线程结束...")
}
