package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("a.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("err:", err)
	}

	// bufio 读取文件
	var fileStr string
	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n') // 表示一次读取一行
		if err == io.EOF {
			fileStr += str
			fmt.Println("读取完毕")
			break
		}
		if err != nil {
			fmt.Println("err:", err)
		}
		fileStr += str
	}
	fmt.Println(fileStr)
}
