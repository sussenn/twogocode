package main

import "fmt"

func test(n1 int) {
	n1 = n1 + 10
	//n1 = 30 = 10 + 20
	fmt.Println("test() n1=", n1)
}

// *int 引用类型
func test01(n2 *int) {
	*n2 = *n2 + 10
	//打印的是引用值 *n2, 而不是地址值n2
	fmt.Println("test01() n2=", *n2)
}

func main() {
	num1 := 20
	//num调用函数后, 传递过去是拷贝值(即值传递),所以在test()内修改了值,并不影响test()外num的值!
	test(num1)
	//num = 20
	fmt.Println("main() num1=", num1)
	//--------------------------------------------------------------------------
	num2 := 2
	//引用传递. 传参过去的是 地址值
	test01(&num2)
	//引用传递, 值被修改了
	fmt.Println("main() num2=", num2)
}
