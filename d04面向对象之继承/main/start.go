package main

import "fmt"

//匿名结构体的使用
type A struct {
	Name string
	age  int
}

func (a *A) sayOK() {
	fmt.Println("A is OK.", a.Name)
}

func (a *A) hello() {
	fmt.Println("hello.", a.Name)
}

type B struct {
	Name  string
	score float64
}

type C struct {
	//匿名结构体
	A
	//有名结构体
	b    B
	Name string
}

func main() {
	var c C
	//常规写法 b.A.xxx
	c.A.Name = "tom"
	//b是有名结构体, 必须声明b
	c.b.Name = "jack"
	c.Name = "mark" //这里Name字段属于结构体B	方法是属于结构体A
	//简写
	c.age = 10
	c.sayOK()
	c.hello()
}
