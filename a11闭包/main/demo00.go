package main

import "fmt"

//AddUpper()是一个闭包. n只初始化一次!多次调用此函数,n会累加!
//AddUpper是函数名, ()是入参, func...是返回值, 入参"x"可省略
//add函数里的匿名func和变量n,str 即构成闭包环境
func AddUpper() func(x int) int {
	var n int = 10
	var str = "hello"
	//将匿名函数func...作返回值return
	return func(x int) int {
		n = n + x
		str += string(36)         //36 即 '$' 对应字码
		fmt.Println("str =", str) //hello$  hello$$	hello$$$
		return n                  //[注意] 这里如果直接return n+x,则n不进行累加!
	}
}

func main() {
	//AddUpper()并没有入参
	//使用变量 f 接收 add()的返回值, 即 f = func(x int) int
	//也就是说 f 此时成为一个带有入参的函数了!
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
}
