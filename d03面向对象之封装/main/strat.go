package main

import (
	"fmt"
	person "mygithub/twogocode/d03面向对象之封装/model"
)

func main() {
	//通过有参构造方法创建对象
	//返回值p是指针类型
	p := person.NewPerson("mary")
	//本质上是写作 (*p).SetAge
	p.SetAge(18)
	p.SetSal(1000.00)

	fmt.Printf("员工姓名:%v\t年龄:%v\t薪资:%v\n", p.Name, p.GetAge(), p.GetSal())
}
