package main

import (
	"fmt"
	"mygithub/twogocode/d01工厂模式解决私有结构体引入/student"
)

func main() {
	//返回值 stu是指针类型
	stu := student.NewStudent("tom", 88.8)
	//"*"取指向数据值
	fmt.Println(*stu)
	//(*stu)可简写为 stu
	//fmt.Println("name=",(*stu).Name,"score=",(*stu).Score)
	//fmt.Println("name=",stu.Name,"score=",stu.Score)

	//当结构体的字段是私有时,使用方法进行获取该字段
	fmt.Println("name=", stu.Name, "score=", stu.GetScore())
}
