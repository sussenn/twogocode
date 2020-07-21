package main

import "fmt"

func test() {
	defer func() {
		//recover()内置函数,用于捕获异常
		if err := recover(); err != nil {
			fmt.Println("捕获到异常,在此处进行处理:", err)
		}
	}() //匿名函数后的"()"代表执行此函数
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("发生异常,并不会执行到此.res=", res)
}

func main() {
	test()
	fmt.Println("后续执行的代码...")
}
