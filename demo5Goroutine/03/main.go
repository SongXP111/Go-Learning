package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 获取当前计算机CPU个数
	cpuNum := runtime.NumCPU()
	fmt.Println("Numbers of CPU: ", cpuNum)

	// 可以自己设置使用多个CPU
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")
}
