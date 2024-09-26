package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./a.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	// 操作文件
	fmt.Println(file)
}
