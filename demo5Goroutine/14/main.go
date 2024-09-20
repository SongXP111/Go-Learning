package main

import (
	"fmt"
	"time"
)

// panic
func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 50)
		fmt.Println("Hello World")
	}
}

func test() {
	defer func() {
		// 捕获test抛出的panic
		if err := recover(); err != nil {
			fmt.Println("抛出异常: ", err)
		}
	}()
	// var myMap map[int]string
	// myMap[0] = "golang" // error
}

func main() {
	go sayHello()
	go test()

	time.Sleep(time.Second)
}
