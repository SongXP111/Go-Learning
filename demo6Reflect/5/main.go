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

// 打印字段
func PrintStructField(s interface{}) {
	// 判断参数是不是结构体类型
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的参数不是结构体")
		return
	}

	// 1. 通过类型变量 Field获取结构体的字段
	field0 := t.Field(0)
	fmt.Println(field0)
	fmt.Println("字段名称: ", field0.Name)
	fmt.Println("字段类型: ", field0.Type)
	fmt.Println("字段标签: ", field0.Tag.Get("json"))
	fmt.Println("字段标签: ", field0.Tag.Get("form"))
	fmt.Println("-------------------------")

	// 2. 通过类型变量 FieldByName
	field1, ok := t.FieldByName("Age")
	if ok {
		fmt.Println("字段名称: ", field1.Name)
		fmt.Println("字段类型: ", field1.Type)
		fmt.Println("字段标签: ", field1.Tag.Get("json"))
	}
	fmt.Println("-------------------------")

	// 3. 通过类型变量 NumField
	var fieldCount = t.NumField()
	fmt.Printf("结构体有%v个属性\n", fieldCount)
	fmt.Println("-------------------------")

	// 4. 通过值变量获取结构体属性对应的值
	fmt.Println(v.FieldByName("Name"))
	fmt.Println(v.FieldByName("Age"))
	fmt.Println(v.FieldByName("Score"))
	fmt.Println("-------------------------")

	// 5. for循环获取属性
	for i := 0; i < fieldCount; i++ {
		fmt.Printf("属性名称: %v, 属性值: %v, 属性类型: %v, 属性tag: %v\n", t.Field(i).Name, v.Field(i), t.Field(i).Type, t.Field(i).Tag.Get("json"))
	}
}

// 打印结构体方法
func PrintStructFn(s interface{}) {
	// 判断参数是不是结构体类型
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的参数不是结构体")
		return
	}

	// 1. 通过类型变量里面的Method获取结构体方法
	method0 := t.Method(0)    // 和结构体的方法没有关系，和ASCII码有关系
	fmt.Println(method0.Name) // GetInfo
	fmt.Println(method0.Type) // func(main.Student) string
	fmt.Println("-------------------------")

	// 2. 通过类型变量获取结构体有多少个方法
	method1, ok := t.MethodByName("Print")
	if ok {
		fmt.Println(method1.Name) // Print
		fmt.Println(method1.Type) // func(main.Student)
	}
	fmt.Println("-------------------------")

	// 3. 通过值变量执行方法
	v.Method(1).Call(nil)
	v.MethodByName("Print").Call(nil)
	info := v.MethodByName("GetInfo").Call(nil)
	fmt.Println(info)
	fmt.Println("-------------------------")

	// 4. 执行方法传入参数（必须传入地址值）
	var params []reflect.Value
	params = append(params, reflect.ValueOf("李四"))
	params = append(params, reflect.ValueOf(23))
	params = append(params, reflect.ValueOf(100))
	v.MethodByName("SetInfo").Call(params)
	info2 := v.MethodByName("GetInfo").Call(nil)
	fmt.Println(info2)
	fmt.Println("-------------------------")

	// 5. 获取方法数量
	fmt.Println("方法数量：", t.NumMethod())
}

func main() {
	stu1 := Student{
		Name:  "张三",
		Age:   18,
		Score: 98,
	}
	//PrintStructField(stu1)
	PrintStructFn(&stu1)
	//fmt.Println(stu1.GetInfo())
}
