package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写锁
// 多个读操作并发，多个写操作完全互斥

var wg sync.WaitGroup
var mutex sync.RWMutex

// 写的方法
func write() {
	mutex.Lock()
	fmt.Println("执行写操作")
	time.Sleep(time.Second * 2)
	mutex.Unlock()
	wg.Done()
}

func read() {
	mutex.RLock()
	fmt.Println("执行读操作")
	time.Sleep(time.Second * 2)
	mutex.RUnlock()
	wg.Done()
}

// 读的方法

func main() {
	// 开启10个协程执行写操作
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	// 开启10个协程执行读操作
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
}
