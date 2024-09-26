package main

import (
	"fmt"
	"os"
)

func main() {
	// 参数二：文件模式
	// 参数三：权限
	file, err := os.OpenFile("a.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer file.Close()

	// 写入文件

	// 1. WriteString
	//for i := 0; i < 10; i++ {
	//	file.WriteString("饮必甘露，食必嘉禾。循循如仪，不罹罗网 - " + strconv.Itoa(i) + "\r\n")
	//}

	// 2. Write
	var str = "饮必甘露，食必嘉禾。循循如仪，不罹罗网"
	file.Write([]byte(str))
}
