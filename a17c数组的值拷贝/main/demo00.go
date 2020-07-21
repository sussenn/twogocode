package main

import "fmt"

//注意:函数的形参必须是 [3]int, 而[]int 是切片类型!
//数组在传递的时候,长度是数据类型的一部分,必须携带!
func test01(arr [3]int) {
	arr[0] = 100
}

//数组的指针传递(引用传递)
func test02(arr *[3]int) {
	(*arr)[0] = 100
}

func main() {
	//在main()栈中, 初始化一个arr数组. 注: [...] 是通过初始化数组的长度而取值.如长度为3,则[...]为[3]
	arr := [...]int{10, 20, 30}
	//在test01()栈中, 数组是值类型,所以会复制出一个arr数组.
	test01(arr)
	//[10 20 30]	值并未改变
	fmt.Println(arr)
	//---------------------------
	test02(&arr)
	//[100 20 30]	值改变
	fmt.Println(arr)
}
