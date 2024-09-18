package main

import "fmt"

type myFn func(int, int) int

// 类型
type myFloat = float64

func structDemo1() {
	var b myFloat = 12.3
	fmt.Printf("%v, %T\n", b, b)
}

// 结构体大写：public
// 结构体小写：private
type Person struct {
	name string
	age  int
	sex  string
}

func structDemo2() {
	// 实例化Person结构体
	var p1 Person
	p1.name = "张三"
	p1.age = 25
	p1.sex = "Male"

	fmt.Printf("%v, %T\n", p1, p1)
	fmt.Printf("%#v, %T\n", p1, p1) // 更全的结构体信息

	var p2 = new(Person) // 指针
	// 在go中，go允许对结构体指针直接访问，*(p2).name = p2.name
	p2.name = "赵四"
	p2.age = 20
	p2.sex = "Male"
	// (*p2).name = "王五"
	fmt.Printf("%v, %T\n", p2, p2)
	fmt.Printf("%#v, %T\n", p2, p2)

	var p3 = &Person{} // 也是指针
	p3.name = "王五"
	p3.age = 29
	p3.sex = "Male"
	fmt.Printf("%v, %T\n", p3, p3)
	fmt.Printf("%#v, %T\n", p3, p3)

	p4 := Person{
		name: "李六",
		age:  19,
		sex:  "Male",
	}
	fmt.Printf("%v, %T\n", p4, p4)
	fmt.Printf("%#v, %T\n", p4, p4)

	// 指针
	p5 := &Person{
		name: "孙七",
		age:  22,
		sex:  "Female",
	}
	fmt.Printf("%v, %T\n", p5, p5)
	fmt.Printf("%#v, %T\n", p5, p5)
}

func structDemo3() {
	var p1 Person
	p1.name = "张三"
	p1.age = 25
	p1.sex = "Male"

	p2 := p1
	p2.name = "李四"

	// 打印的结果不一样，说明go中的结构体struct是值类型
	fmt.Printf("%#v, %T\n", p1, p1)
	fmt.Printf("%#v, %T\n", p2, p2)
}

type Person2 struct {
	name   string
	age    int
	sex    string
	height int
}

func (p Person2) PrintInfo() {
	fmt.Printf("Name: %v\n", p.name)
	fmt.Printf("Age: %v\n", p.age)
}

// 修改必须用结构体指针，因为结构体是值类型
func (p *Person2) SetInfo(name string, age int) {
	p.name = name
	p.age = age
}

func structDemo4() {
	p1 := Person2{
		name:   "张三",
		age:    20,
		sex:    "Male",
		height: 188,
	}
	p1.PrintInfo()

	p2 := Person2{
		name:   "李四",
		age:    30,
		sex:    "Male",
		height: 175,
	}
	p2.PrintInfo()
	p2.SetInfo("王五", 22) // p2的属性被修改了
	p2.PrintInfo()
}

type MyInt int

func (m MyInt) printInfo() {
	fmt.Println("我是自定义MyInt类型的自定义方法")
}

func structDemo5() {
	var a MyInt = 20
	a.printInfo()
}

// 结构体匿名字段，字段名称不能重复
type Person3 struct {
	string
	int
}

func structDemo6() {
	p := Person3{
		"张三",
		20,
	}
	fmt.Println(p)
}

// 结构体字段可以是基本数据类型，也可以是map，切片，结构体
type Person4 struct {
	Name    string
	Age     int
	Hobby   []string
	InfoMap map[string]string
}

func structDemo7() {
	var p Person4
	p.Name = "张三"
	p.Age = 20
	p.Hobby = make([]string, 3, 6)
	p.Hobby[0] = "Coding"
	p.Hobby[1] = "Sleeping"
	p.Hobby[2] = "Singing"
	p.InfoMap = make(map[string]string)
	p.InfoMap["address"] = "US"
	p.InfoMap["phone"] = "22112211"

	fmt.Printf("Name: %v\n", p.Name)
	fmt.Printf("Age: %v\n", p.Age)
	fmt.Printf("Hobby: %v\n", p.Hobby)
	fmt.Printf("Info: %v\n", p.InfoMap)
}

type User struct {
	Username string
	Password string
	Address  Address
}

type Address struct {
	Name  string
	Phone string
	City  string
}

func structDemo8() {
	var u User
	u.Username = "Xipeng"
	u.Password = "123456"
	// u.Address.Name = "全运村"
	u.Address = Address{
		"全运村",
		"123456789",
		"天津",
	}

	fmt.Printf("Username: %v\n", u.Username)
	fmt.Printf("Password: %v\n", u.Password)
	fmt.Printf("Address: %v\n", u.Address)
}

// 继承
// 父亲结构体
type Animal struct {
	Name string
}

func (a Animal) run() {
	fmt.Printf("动物%v在跑\n", a.Name)
}

type Dog struct {
	// Name   string
	Age    int
	Animal // 结构体嵌套，也是继承
	// *Animal // 也可以这样写
}

func (d Dog) wang() {
	fmt.Printf("小狗%v在汪汪叫\n", d.Name)
}

func (d Dog) run() {
	fmt.Printf("小狗%v在跑\n", d.Name)
}

func structDemo9() {
	var d = Dog{
		// Name: "阿黄",
		Age: 3,
		Animal: Animal{
			Name: "阿奇",
		},
		// Animal: &Animal{
		// 	Name: "阿奇",
		// },
	}

	// 如果Dog有run方法，用Dog的run，如果没有，用Animal的run
	d.run()
	d.wang()
}

// json
