package main

import "fmt"

//接口的"实例化"	接口不能创建实例! 但可以指向实现了该接口方法的对象, 即将man赋值给接口
type MyInterface interface {
	//接口内部不得有变量!
	Say()
	Eat()
}

type Man struct {
	Name string
}

//Man结构体实现了 接口方法
func (m Man) Say() {
	fmt.Println("hello world")
}
func (m Man) Eat() {
	fmt.Println("chichichi")
}

func main() {
	var man Man
	//只要实现了接口所有方法,那么该对象即可赋值给此接口.		(struct类型, 自定义数据类型都行)
	//man实现了Say(), 所以可以赋值给MyInterface接口的实例!
	var myIn MyInterface = man
	myIn.Say()
}
