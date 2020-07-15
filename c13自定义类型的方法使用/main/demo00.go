package main

import "fmt"

//给自定义类型 创建方法
type myInt int

func (i myInt) seaHello() {
	fmt.Println("hello")
}

func (i *myInt) change() {
	*i += 1
}

//======================================================
type Student struct {
	Name string
	Age  int
}

//给结构体实现String方法	类似Java toString()
func (s *Student) String() string {
	str := fmt.Sprintf("Name=[%v] Age=[%v]", s.Name, s.Age)
	return str
}

func main() {
	var n myInt = 10
	n.seaHello()
	n.change()
	fmt.Println("n值已改变:n=", n)

	//======================================================
	stu := Student{"tom", 18}
	//Name=[tom] Age=[18]
	fmt.Println(&stu)
}
