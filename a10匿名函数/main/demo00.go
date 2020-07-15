package main

import "fmt"

func main() {
	//匿名函数使用1:
	//只能调用一次. 直接传参,然后使用变量res1接收结果
	//格式: 变量 = func(形参){逻辑}(入参)
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("res1 =", res1)

	//------------------------------------
	//匿名函数使用2:
	//可调用多次. 先不传参,使用变量接收函数. 然后调用变量a传参
	a := func(n1 int, n2 int) int {
		return n1 + n2
	}
	res2 := a(1, 2)
	res3 := a(2, 3)
	fmt.Println("res2 =", res2)
	fmt.Println("res3 =", res3)

	//-----------------------------------------
	//全局匿名函数
	res4 := fun1(100, 200)
	fmt.Println("res4 =", res4)

}

//全局匿名函数
var (
	fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)
