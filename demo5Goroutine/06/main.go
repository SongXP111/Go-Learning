package main

import (
	"fmt"
	"sync"
	"time"
)

// 需求：要统计1-120000的数字中哪些是素数？
// 用for循环实现
// 利用goroutine
// 用四个协程统计，每个统计30000个数字
// start = (n - 1) * 30000 + 1, end = n * 30000

var wg sync.WaitGroup

func test(n int) {
	start := (n-1)*30000 + 1
	end := n * 30000
	for num := start; num < end; num++ {
		var flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			// fmt.Println(num, "是素数")
		}
	}
	wg.Done()

}

func main() {
	start := time.Now().UnixMilli()
	for i := 1; i < 4; i++ {
		wg.Add(1)
		go test(i)
	}
	wg.Wait()
	end := time.Now().UnixMilli()
	fmt.Println(end - start) // 204, 192, 197, 193
}
