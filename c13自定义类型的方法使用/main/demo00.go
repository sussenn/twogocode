package main

import "fmt"

//给自定义类型 创建方法
type myInt int

func (i myInt) seaHello() {
	fmt.Println("hello")
}

//传参myInt类型, 返回值也是
func (i *myInt) change() myInt {
	*i += 1
	return *i
}

//======================================================
type Student struct {
	Name string
	Age  int
}

//给结构体实现String方法,当使用"&"时,默认会调用String()	//类似Java toString()
func (s *Student) String() string {
	return fmt.Sprintf("{Name=%v,Age=%v}", s.Name, s.Age)
}

func main() {
	var n myInt = 10
	n.seaHello()
	//因为 *myInt是指针类型,所以应该使用&n进行调用.	但是底层作过优化,可以不用写"&"
	//res := (&n).change()
	res := n.change()
	fmt.Println("n值已改变:n=", n)
	fmt.Println("返回值是myint类型:res=", res)

	//======================================================
	//student := Student{Name: "tom",Age: 18}	//2种写法
	stu := Student{"tom", 18}
	//{Name=tom,Age=18}
	fmt.Println(&stu)
}
