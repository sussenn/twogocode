package main

import "fmt"

//[注意] 如果有引入其他包 如utils,
//则执行顺序是: utils包的全局变量 --> utils包的init() --> main包的全局变量 --> main包的init() --> main包的main()
//------------------------------------
//init() 先于main()执行
func init() {
	fmt.Println("我是 init()")
}

//比init()更先执行的是 全局变量(包括 函数变量!)
var fun = test()

func test() int {
	fmt.Println("我是全局变量 (全局函数变量) 我最快")
	return 100
}

func main() {
	fmt.Println("我是main()")
}
