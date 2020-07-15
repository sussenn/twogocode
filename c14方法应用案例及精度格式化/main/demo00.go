package main

import "fmt"

type Calculator struct {
	Num1 float64
	Num2 float64
}

func (ca *Calculator) getRes(operator byte) float64 {
	res := 0.0
	switch operator {
	case '+':
		res = ca.Num1 + ca.Num2
	case '-':
		res = ca.Num1 - ca.Num2
	case '*':
		res = ca.Num1 * ca.Num2
	case '/':
		res = ca.Num1 / ca.Num2
	default:
		fmt.Println("运算符输入有误!")
	}
	return res
}

func main() {

	var ca Calculator
	ca.Num1 = 10.15
	ca.Num2 = 20.28
	//方法getRes 使用的是指针类型*Calculator, 所以调用该方法,要用地址值"&"
	//res := (&ca).getRes('/')
	//底层优化后,可省略"&"
	res := ca.getRes('/')
	//fmt.Sprintf("%.6f", res)	//数据精度的格式化
	fmt.Println("计算结果:", fmt.Sprintf("%.6f", res))
}
