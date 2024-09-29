package main

import (
	"fmt"
	"os"
)

// 编写一个函数，接收两个文件路径

func copyFile(scrFileName, destFileName string) error {
	input, err := os.ReadFile(scrFileName)
	if err != nil {
		fmt.Println("read file err:", err)
		return err
	}
	err = os.WriteFile(destFileName, input, 0644)
	if err != nil {
		fmt.Println("write file err:", err)
		return err
	}
	fmt.Println("复制文件成功")
	return nil
}

func main() {
	src := "/Users/xipengsong/Desktop/github/Go/demo7File/7/a.txt"
	dst := "/Users/xipengsong/Desktop/github/Go/demo7File/8/a.txt"
	err := copyFile(src, dst)
	if err != nil {
		fmt.Println("copy file err:", err)
		return
	}
}
