package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 向intChan放入120000个数
func putNum(intChan chan int) {
	for i := 2; i < 120; i++ {
		intChan <- i
	}
	close(intChan)
	wg.Done()
}

// 从intChan取出数据，并判断是否为素数，如果是，就把得到的素数放在primeChan
func primeNum(intChan, primeChan chan int, exitChan chan bool) {
	for num := range intChan {
		var flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			// fmt.Println(num, "是素数")
			primeChan <- num
		}
	}
	// 给exitChan传入一条数据
	exitChan <- true
	wg.Done()
}

// printPrime打印素数
func printPrime(primeChan chan int) {
	for v := range primeChan {
		fmt.Println(v, "是素数")
	}
	wg.Done()
}

func main() {
	start := time.Now().UnixMilli()

	intChan := make(chan int, 50000)
	primeChan := make(chan int, 50000)
	exitChan := make(chan bool, 16) // 标识primeChan close

	// 存放数字的协程
	wg.Add(1)
	go putNum(intChan)

	// 统计数字的协程
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go primeNum(intChan, primeChan, exitChan)
	}

	// 打印素数的协程
	wg.Add(1)
	go printPrime(primeChan)

	// 判断exitChan是否存满值
	wg.Add(1)
	go func() {
		for i := 0; i < 16; i++ {
			<-exitChan
		}
		// 关闭primeChan
		close(primeChan)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("执行完毕")

	end := time.Now().UnixMilli()
	fmt.Println(end - start) // 113, 110, 113
}
