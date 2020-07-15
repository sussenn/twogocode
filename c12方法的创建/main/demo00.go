package main

import "fmt"

type Man struct {
	Name string
	Age  int
	Num  int
}

//(m Man) 声明"方法"的拥有者.	除了struct,也可以是自定义类型
func (m Man) eat() {
	fmt.Println("通过方法传参Name=:", m.Name)
}

func (m Man) getSum(n int, s int) int {
	return n + s
}

//指针类型
func (m *Man) handle() int {
	m.Num = 100          //Num是int类型!
	return m.Age + m.Num //返回值是int类型
}

func main() {
	m := Man{}
	m.Name = "sussen"
	//通过对象,调用方法
	m.eat()

	//====================================================
	//[注意: 实际上,调用方法getSum() 把 m 也作为参数传递给方法了
	sum := m.getSum(1, 2)
	fmt.Println("传参调用方法获取返回值:", sum)

	//====================================================
	m.Age = 10
	m.Num = 10
	//调用方法后,因为是指针类型,所以Num值能被修改
	res := m.handle()
	fmt.Println("指针类型结构体调用方法:", res)
}
