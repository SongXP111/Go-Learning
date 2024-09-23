package main

import (
	"fmt"
	"reflect"
)

// 结构体反射
type Student struct {
	Name  string `json:"name" form:"username"` // 标签可以设置多个
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func (s Student) GetInfo() string {
	var str = fmt.Sprintf("姓名: %v, 年龄: %v, 成绩: %v", s.Name, s.Age, s.Score)
	return str
}

func (s *Student) SetInfo(name string, age int, score int) {
	s.Name = name
	s.Age = age
	s.Score = score
}

func (s Student) Print() {
	fmt.Println("这是一个打印方法")
}

func reflectChangeStruct(s interface{}) {
	// 判断参数是不是结构体类型
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() != reflect.Ptr {
		fmt.Println("传入的参数不是结构体指针类型")
		return
	} else if t.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的参数不是结构体指针类型")
		return
	}

	// 修改结构体属性的值
	name := v.Elem().FieldByName("Name")
	name.SetString("李四")

	age := v.Elem().FieldByName("Age")
	age.SetInt(23)

	score := v.Elem().FieldByName("Score")
	score.SetInt(100)

}

func main() {
	stu1 := Student{
		Name:  "张三",
		Age:   18,
		Score: 98,
	}
	fmt.Println(stu1.GetInfo())
	reflectChangeStruct(&stu1)
	fmt.Println(stu1.GetInfo())
}
