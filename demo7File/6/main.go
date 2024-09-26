package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("a.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer file.Close()

	// 通过buffer写入
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("饮必甘露，食必嘉禾。循循如仪，不罹罗网\n") // 将数据存入缓存

	}
	writer.Flush() // 将缓存数据写入文件
}
