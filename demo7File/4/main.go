package main

import (
	"fmt"
	"os"
)

func main() {
	// ioutil
	//byteStr, err := ioutil.ReadFile("a.txt")
	byteStr, err := os.ReadFile("a.txt")
	if err != nil {
		fmt.Println("read file err:", err)
		return
	}
	fmt.Println(string(byteStr))
}
