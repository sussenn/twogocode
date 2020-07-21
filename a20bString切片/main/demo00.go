package main

import "fmt"

//string 切片使用
func main() {
	//string底层是[]byte 所以可以切片
	str := "helloworld"
	//slice[0]地址值指向 str[5]
	slice := str[5:]
	fmt.Println("string切片:", slice)
	//--------------------------------------------------
	//str[0] = 'a'		//编译报错.	[string是不可变的]
	//方式一: 可以将string强转[]byte后, 重新赋值	(对中文无效)
	bytes := []byte(str)
	bytes[0] = 'a' //不能赋值中文!
	//再强转回string
	str = string(bytes)
	fmt.Println("强转后,重新赋值:", str)
	//-------------------------------------------------
	//方式二: 使用[]rune 处理中文字符
	arr := []rune(str)
	arr[0] = '大'
	str = string(arr)
	fmt.Println("针对中文字符,强转后,重新赋值:", str)
}
