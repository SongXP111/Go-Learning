package main

import "fmt"

type Address struct {
	Name  string
	Phone int
}

// Golang中空接口和类型断言使用细节
func main() {
	var userinfo = make(map[string]interface{})
	userinfo["username"] = "张三"
	userinfo["age"] = 20
	userinfo["hobby"] = []string{"睡觉", "吃饭"}

	fmt.Println(userinfo["age"])
	fmt.Println(userinfo["hobby"])

	// fmt.Println(userinfo["hobby"][0]) // 报错

	var address = Address{
		Name:  "李四",
		Phone: 15544446666,
	}
	userinfo["address"] = address
	fmt.Println(address.Name)        // 李四
	fmt.Println(userinfo["address"]) // {李四 15544446666}

	// var name = userinfo["address"].Name // 报错

	// 需要使用类型断言
	hobby, ok := userinfo["hobby"].([]string)
	if ok {
		fmt.Println(hobby[0]) // 可以用索引获取值
	}

	addr, ok := userinfo["address"].(Address)
	if ok {
		fmt.Println(addr.Name) // 可以获取结构体的属性
	}
}
