package main

import "fmt"

// 单向管道
func main() {
	// gin1. 默认情况是双向管道
	ch1 := make(chan int, 2)
	ch1 <- 10
	ch1 <- 12
	m1 := <-ch1
	m2 := <-ch1
	fmt.Println(m1, m2)

	// 2. 管道声明为只写
	ch2 := make(chan<- int, 2)
	ch2 <- 10
	ch2 <- 12
	// <-ch2 // ./main.go:19:4: invalid operation: cannot receive from send-only channel ch2 (variable of type chan<- int)

	// 3. 管道声明为只读
	// ch3 := make(<-chan int, 2)
	// ch3 <- 23 // ./main.go:23:2: invalid operation: cannot send to receive-only channel ch3 (variable of type <-chan int)
}
