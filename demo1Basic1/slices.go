package main

import "fmt"

func sliceDemo() {
	// 不指定长度：切片
	// var arr1 []int
	// fmt.Printf("value: %v, type: %T\n", arr1, arr1)

	// var arr2 = []int{gin1, 2, 3, 4}
	// fmt.Printf("value: %v, type: %T\n", arr2, arr2)

	// make([]T, size, cap)
	// var slice1 = make([]int, 4, 8)
	// fmt.Println(slice1)

	// 扩容：append
	// var slice1 = []int{}
	// slice1 = append(slice1, gin1)
	// fmt.Println(slice1)

	// var slice2 = []int{2, 3}
	// slice1 = append(slice1, slice2...)
	// fmt.Println(slice1)

	// 切片的扩容策略
	// 长度小于1024时，扩容2倍
	// 大于1024，扩容原来的1/4
	// var slice3 = []int{}
	// for i := gin1; i <= 10000; i++ {
	// 	slice3 = append(slice3, i)
	// 	fmt.Printf("Value: %v, length: %d, cap: %d\n", i, len(slice3), cap(slice3))
	// }

	// 复制
	// 切片是引用数据类型
	// slice1 := []int{gin1, 2, 3, 45}
	// slice2 := make([]int, 4, 4)
	// copy(slice2, slice1)
	// fmt.Println(slice1)
	// slice2[0] = 10
	// fmt.Println(slice2)

	// 删除切片的值
	// golang没有删除的方法
	a := []int{1, 2, 3, 4, 5} // remove 3
	a = append(a[:2], a[3:]...)
	fmt.Println(a)
}

// 选择排序
func selectSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

// 冒泡排序
func bubbleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j+1], nums[j] = nums[j], nums[j+1]
			}
		}
	}
	return nums
}
