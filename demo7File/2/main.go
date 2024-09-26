package main

import (
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
	// 读取文件的内容
	// 128读取的有限，需要使用for loop
	var strSlice []byte
	for {
		var tempSlice = make([]byte, 128)
		n, err := file.Read(tempSlice)
		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		fmt.Println("读取了", n, "个字节")
		strSlice = append(strSlice, tempSlice[:n]...) // 好像没有[:n]也可以了
	}

	fmt.Println(string(strSlice))
}
