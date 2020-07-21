package main

import "fmt"

//内置函数: go内部提供,可直接使用的函数
func main() {
	//1. len() 获取长度
	//2. new() 分配内存,分配值类型会返回指针
	num1 := 100
	fmt.Printf("num1类型=%T, 值=%v, 地址值=%v \n", num1, num1, &num1)

	num2 := new(int) //返回指针类型
	*num2 = 100
	//注意num2的值 和指向值
	//num2类型=*int, 值=0xc00000a100, 地址值=0xc000006030, 指针指向值=100
	//num2 -> 地址值0xc000006030 -> (值0xc00000a100 -> 指针地址值0xc00000a100) -> 指针值100
	fmt.Printf("num2类型=%T, 值=%v, 地址值=%v, 指针指向的值=%v", num2, num2, &num2, *num2)

	//3. make()分配内存,用于分配引用类型
	//...
}
