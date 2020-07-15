package main

import (
	"fmt"
	"strings"
)

//闭包相当于 "类", 类中方法和变量的操作即构成一个环境--"闭包"
//闭包使用演示
func makeSuffix(suffix string) func(name string) string {
	return func(name string) string {
		//如果 传入文件没有后缀,则给拼接后缀
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		//有后缀,则直接返回
		return name
	}
}

func main() {
	f2 := makeSuffix(".jpg")

	fmt.Println("没有后缀的文件处理:", f2("hello"))
	fmt.Println("有后缀的文件处理:", f2("world.jpg"))
}
