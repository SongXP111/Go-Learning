package main

import "fmt"

func arrDemo() {
	// var arr1 [3]int
	// var arr2 [4]int
	// var strArr [3]string

	// fmt.Printf("arr1:%T, arr2:%T, arr3:%T", arr1, arr2, strArr)

	// 初始化
	// var arr = []string{"java", "python", "golang"}
	// var arr [3]string
	// arr[0] = "java"
	// arr[1] = "python"
	// arr[2] = "golang"
	// fmt.Println(arr)

	// 数组：值类型
	var arr1 = [...]int{1, 2, 3}
	arr2 := arr1
	arr1[0] = 10
	fmt.Println(arr1, arr2)

	// 切片：引用类型
	// 类似于指针
	var arr3 = []int{1, 2, 3}
	arr4 := arr3
	arr3[0] = 10
	fmt.Println(arr3, arr4)

}

func twoDArrayDemo() {
	var arr = [3][2]string{
		{
			"a", "b",
		},
		{
			"c", "d",
		},
		{
			"e", "f",
		},
	}

	fmt.Println(arr)

	// for _, v1 := range arr {
	// 	for _, v2 := range v1 {
	// 		fmt.Println(v2)
	// 	}
	// }

	// for i := 0; i < len(arr); i++ {
	// 	for j := 0; j < len(arr[0]); j++ {
	// 		fmt.Println(arr[i][j])
	// 	}
	// }
}
