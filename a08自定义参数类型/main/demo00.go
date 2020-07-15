package main

import "fmt"

//使用函数作 自定义类型, 作形参
type myFunType func(int, int) int

//此函数 作入参
func getSum(n1 int, n2 int) int {
	return n1 + n2
}

//最终进行调用的函数
func myFun(ge myFunType, n1 int, n2 int) int {
	return ge(n1, n2)
}

func main() {
	res := myFun(getSum, 10, 20)
	//res = 30
	fmt.Println("res=", res)
}
