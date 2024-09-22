package main

import (
	"fmt"
	"sync"
	"time"
)

var count = 0
var wg sync.WaitGroup

// 互斥锁
var mutex sync.Mutex

func test() {
	// 加锁
	mutex.Lock()

	count++
	fmt.Println("the count is: ", count)
	time.Sleep(time.Millisecond)

	// 解锁
	mutex.Unlock()

	wg.Done()
}

func main() {
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go test()
	}
	wg.Wait()
}
