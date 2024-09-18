package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

// 可变参数
// 可输入N个数
func sumN(x ...int) int {
	// fmt.Printf("Value: %v, Type: %T", x, x)
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

// 可变参数放后边
func sumN2(x int, y ...int) int {
	// fmt.Printf("Value: %v, Type: %T", x, x)
	sum := x
	for _, v := range y {
		sum += v
	}
	return sum
}

// return一次可以返回多个值
func calc0(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

// 给返回值命名
func calc1(x, y int) (sum, sub int) {
	// fmt.Println(sum, sub)
	sum = x + y
	sub = x - y
	return
}

// 排序
// int类型升序排序
func sortIntAsc(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

// int类型升序排序
func sortIntDesc(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

// type calc func(int, int) int
type calcType func(int, int) int

// 函数作为另一个函数的参数
func calc(x, y int, cb calcType) int {
	return cb(x, y)
}

// 函数作为返回值
func do(o string) calcType {
	switch o {
	case "+":
		return add
	case "-":
		return sub
	case "*":
		return func(x, y int) int {
			return x * y
		}
	default:
		return nil
	}
}

// 匿名函数
func afunc() {
	// 匿名自执行函数
	func() {
		fmt.Println("我是匿名函数")
	}()

	// 接受参数
	func(x, y int) {
		fmt.Println(x * y)
	}(10, 20)

	// 存在var的匿名函数
	var fn = func(x, y int) int {
		return x * y
	}

	fmt.Println(fn(2, 3))
}

// 递归
func fn1(n int) int {
	if n > 1 {
		return n + fn1(n-1)
	} else {
		return 1
	}
}

func fn3(n int) int {
	if n > 1 {
		return n * fn3(n-1)
	} else {
		return 1
	}
}
