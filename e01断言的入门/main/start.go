package main

import "fmt"

//由于接口是一般类型，不知道具体类型，如果要转成具体类型，就需要使用类型断言
//类型Java 类型强转
type Point struct {
	x int
	y int
}

func main() {
	//所有类型都实现了 空接口. 即任何变量都可以赋值给空接口
	var a interface{}
	var point = Point{1, 2}
	//将对象point赋值给 空接口变量a
	a = point
	fmt.Println("空接口变量a接收对象point后:", a)
	var b Point
	//将空接口对象a赋值给 对象b	//使用断言
	b = a.(Point)
	fmt.Println("对象b接收空接口变量a后:", b)

	//=============================================================
	var n interface{}
	var f float64 = 2.1
	n = f
	//空接口类型变量n 强转为 float64类型的t
	if t, ok := n.(float64); ok {
		fmt.Printf("t的类型:%T, 值:%v\n", t, t)
	} else {
		fmt.Printf("强转失败...")
	}

}
