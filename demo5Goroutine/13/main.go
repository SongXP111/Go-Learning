package main

import (
	"fmt"
	"time"
)

// select多路复用
func main() {
	// gin1. 10个int的管道
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	// 2. 5个string的管道
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	// 使用select获取channel里面的数据不需要关闭channel
	for {
		select {
		case v := <-intChan:
			fmt.Printf("从intChan里读取数据:%d\n", v)
			time.Sleep(time.Millisecond * 50)
		case v := <-stringChan:
			fmt.Printf("从stringChan里读取数据:%v\n", v)
			time.Sleep(time.Millisecond * 50)
		default:
			fmt.Println("数据获取完毕")
			return // 注意退出
		}
	}
}
