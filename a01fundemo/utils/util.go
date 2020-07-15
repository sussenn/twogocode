package utils

import "fmt"

//函数名首字目大写, 即public.
func Cal(n1 float64, n2 float64, ope byte) float64 {
	var res float64
	switch ope {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("请输入正确运算符")
	}
	return res
}
