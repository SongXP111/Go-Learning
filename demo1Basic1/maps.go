package main

import (
	"fmt"
)

func mapDemo1() {
	// 创建
	// var userinfo = make(map[string]string)
	// var userinfo = map[string]string{}
	userinfo := map[string]string{}

	userinfo["username"] = "Xipeng"
	userinfo["age"] = "24"
	userinfo["gender"] = "Male"
	fmt.Println(userinfo)

	// 遍历
	for k, v := range userinfo {
		fmt.Printf("Key: %v, Value: %v\n", k, v)
	}

	// 查找
	// value, ok := userinfo["username"]

	// 删除
	delete(userinfo, "gender")
	fmt.Println(userinfo)
}

func mapDemo2() {
	// 元素为map类型的切片
	// var userinfo = []string{}
	// var userinfo = make([]map[string]string, 3, 3)
	// fmt.Println(userinfo[0]) // map默认值为nil

	// 值为切片类型的map
	// var userinfo = make(map[string][]string)

	// map是引用数据类型
}
