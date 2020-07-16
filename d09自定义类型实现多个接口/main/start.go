package main

import "fmt"

//将多个接口赋值给一个对象	即一个对象实现多个接口的方法
type Ainter interface {
	Test01()
}
type Binter interface {
	Test02()
}

//一个接口里包含多个接口
type CCC interface {
	Ainter
	Binter
	Test03()
}

type Man struct {
}

//Man对象实现A,B两个接口的方法
func (m Man) Test01() {
	fmt.Println("Man A Test01()")
}
func (m Man) Test02() {
	fmt.Println("Man B Test02()")
}

//Man要想实现CCC接口,必须先将CCC接口里的A,B接口方法都实现
func (m Man) Test03() {
	fmt.Println("Man CCC Test03()")
}

func main() {
	var man Man
	var ainter Ainter = man
	var binter Binter = man
	ainter.Test01()
	binter.Test02()

	//===============================
	var man2 Man
	//Man实现了CCC接口内所有的方法,才能使用CCC接口
	var c CCC = man2
	c.Test01()
	c.Test02()
	c.Test03()
}
