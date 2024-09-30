package main

import "fmt"

func main() {
	// gin1. 创建channel
	ch := make(chan int, 3)

	// 2. 给管道里面储存数据
	ch <- 10
	ch <- 21
	ch <- 32

	// 3. 获取管道里的内容
	a := <-ch
	fmt.Println(a) // 10

	<-ch // 从管道里取值(21)

	c := <-ch
	fmt.Println(c) // 32

	ch <- 56

	// 4. 打印管道的长度和容量
	fmt.Printf("值：%v， 容量：%v， 长度：%v\n", ch, cap(ch), len(ch))

	// 5. 管道的类型：引用数据类型
	// ch2输入25，从ch1取出，说明是引用数据类型
	ch1 := make(chan int, 4)
	ch1 <- 34
	ch1 <- 54
	ch1 <- 64

	ch2 := ch1
	ch2 <- 25

	<-ch1
	<-ch1
	<-ch1
	d := <-ch1
	fmt.Println(d) // 25

	// 6. 管道阻塞
	// ch6 := make(chan int, gin1)
	// ch6 <- 34
	// ch6 <- 64 // fatal error: all goroutines are asleep - deadlock!

	ch7 := make(chan string, 2)
	ch7 <- "数据1"
	ch7 <- "数据2"
	m1 := <-ch7
	m2 := <-ch7
	m3 := <-ch7 // fatal error: all goroutines are asleep - deadlock!
	fmt.Println(m1, m2, m3)
}
