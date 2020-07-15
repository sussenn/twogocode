package main

import "fmt"

func main() {
	var a = 10
	//var b *int
	var b *int = new(int)
	//b指向a的地址值		//b是指针类型,即存储地址值
	b = &a
	//将b所指向的数据值改变
	*b = 100
	fmt.Println("a=", a)
	//"&"是取本身地址
	fmt.Printf("b本身的地址=%p\n", &b)

	//什么都不写,即默认取指向值的地址值,配合"%p"	因为其是指针类型,输出的是地址值
	//"*"是取所指向的数据值
	fmt.Printf("b所指向的地址=%p\nb所指向的数据值=%v\n", b, *b)
}
