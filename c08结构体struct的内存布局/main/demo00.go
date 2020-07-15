package main

import (
	"fmt"
)

type Man struct {
	Name string
	Age  int
}

//struct是值类型		["."的优先级比"*"高,所以要用"()"]
func main() {
	m1 := Man{}
	m1.Name = "sussen"
	m1.Age = 10
	//m2是地址值,*m2是数据值
	var m2 *Man = &m1 //将m1地址值存入m2空间
	fmt.Println("指针类型的m2,Age=", (*m2).Age)
	m2.Name = "tom"
	//m1.Name=tom	 m2.Name=tom	m2是指针类型,将引用对象的值也修改了
	fmt.Printf("m1.Name=%v\t m2.Name=%v\n", m1.Name, m2.Name)
	//m1的地址值=0xc0000044a0
	fmt.Printf("m1的地址值=%p\n", &m1)
	//m2的地址值=0xc000006028	 m2的值=0xc0000044a0
	fmt.Printf("m2的地址值=%p\t m2存放的值=%p\n", &m2, m2) //[注意]此写法,%p获取m2地址值, %v获取*m2数据值
	//说明: m2的地址值是新开辟的内存(0xc000006028), 内部存放着m1的地址值(0xc0000044a0),指向m1的数值
}
