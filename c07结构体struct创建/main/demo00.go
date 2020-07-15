package main

import "fmt"

//结构体是值类型
type Person struct {
	Name string
	Age  int
	arr  [3]int //数组	值类型
	per  *int   //指针	指针类型
	sli  []int  //切片	指针类型
	mapp map[string]string
}

type Man struct {
	Name string
	Age  int
}

func main() {
	p1 := Person{}
	p1.arr = [3]int{1, 2, 3}
	//指针字段----待定
	p1.per = new(int)
	var a int = 10
	//p1.per指向a地址值
	p1.per = &a
	//将p1.per所指向的数据值改变
	*(p1.per) = 100
	fmt.Println("p1.per本身地址值=", &(p1.per))
	fmt.Println("p1.per所指向的地址值=", p1.per)
	fmt.Println("p1.per所指向的数据值=", *(p1.per))
	//切片字段
	p1.sli = make([]int, 5)
	p1.sli[0] = 100
	//map字段
	p1.mapp = make(map[string]string)
	p1.mapp["技能"] = "吃"
	fmt.Println("复杂类型的结构体字段赋值,p1=", p1)
	//=======================================================
	//结构体的创建方式
	//1 仅声明------------------------------
	var m1 = Man{}
	fmt.Println("struct创建1-仅声明:", m1)
	//2 创建并初始化-------------------------
	m2 := Man{"张三", 18}
	fmt.Println("struct创建2-创建并初始化:", m2)
	//3 创建指针类型, 2种字段赋值方式都可以使用---
	var m3 *Man = new(Man)
	(*m3).Name = "李四"
	m3.Age = 28
	fmt.Println("struct创建3-创建指针类型:", *m3)
	//4 创建指针类型--------------------------
	var m4 *Man = &Man{}
	(*m4).Name = "王五"
	m4.Age = 38
	fmt.Println("struct创建4-创建指针类型:", *m4)
}
