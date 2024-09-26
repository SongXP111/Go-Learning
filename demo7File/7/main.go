package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	str := "饮必甘露，食必嘉禾。循循如仪，不罹罗网"
	err := ioutil.WriteFile("a.txt", []byte(str), 0666) // 现在用os
	if err != nil {
		fmt.Println(err)
		return
	}

}
