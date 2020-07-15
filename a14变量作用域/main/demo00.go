package main

import "fmt"

var (
	age  int    = 10
	Name string = "sussenn"
	//Add := "xxx"
	//此变量定义是错误的!编译不通过. 因为该语句等价于
	//var Add string
	//Add = "xxx" 		//[函数外不得进行变量赋值!]

)

func main() {
	fmt.Println(age)
	fmt.Println(Name)

	//-------------------------------------------
	for i := 0; i < 5; i++ {
		//变量i的作用域 只在此for循环体内有效
		fmt.Println("i =", i)
	}
	//可以将变量在for外定义
	var j int
	for j = 0; j < 5; j++ {
		fmt.Println("内j=", j)
	}
	fmt.Println("j=", j) //j=5; j++到5即跳出循环

	//九九乘法表
	rintMulti(7)
}

//九九乘法
func rintMulti(num int) {
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v \t", j, i, j*i)
		}
		fmt.Println()
	}
}
