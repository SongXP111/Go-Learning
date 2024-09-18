package main

import "fmt"

// 循环遍历管道数据
// for range
func main() {
	// 输入数据
	var ch1 = make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch1 <- i
	}
	close(ch1) // 关闭管道

	// for range循环遍历管道的值
	// 注意：管道没有key
	for v := range ch1 {
		fmt.Println(v)
	}
}
