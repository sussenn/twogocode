package main

import "fmt"

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

//---------------------
//函数类型的parfun 作参数. 即函数作为形参
func myFun(parfun func(int, int) int, num1 int, num2 int) int {
	return parfun(num1, num2)
}

func main() {
	//将getSum()函数 赋值给变量a, a即函数类型
	//a 的值即getSum()返回值
	a := getSum
	//等价于 res := getSum(1,2)
	res := a(1, 2)
	fmt.Println("res=", res)

	//-----------------------------------
	//函数作参数传递时,需省略括号
	res2 := myFun(getSum, 10, 20)
	fmt.Println("res2=", res2)
}
