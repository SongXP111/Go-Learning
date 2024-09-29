package main

import (
	"fmt"
	"io"
	"os"
)

// 编写一个函数，接收两个文件路径

func copyFile(scrFileName, destFileName string) error {
	src, err1 := os.Open(scrFileName)
	dst, err2 := os.OpenFile(destFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	defer src.Close()
	defer dst.Close()

	var tempSlice = make([]byte, 1280)
	for {
		n1, err := src.Read(tempSlice)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if _, err := dst.Write(tempSlice[:n1]); err != nil {
			return err
		}
	}
	fmt.Println("复制文件完毕")
	return nil
}

func main() {
	src := "/Users/xipengsong/Desktop/github/Go/demo7File/2/a.txt"
	dst := "/Users/xipengsong/Desktop/github/Go/demo7File/9/a.txt"
	err := copyFile(src, dst)
	if err != nil {
		fmt.Println("copy file err:", err)
		return
	}
}
