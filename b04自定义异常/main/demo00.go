package main

import (
	"errors"
	"fmt"
)

//在函数中 进行自定义异常,并返回
func myHander(name string) (err error) {
	//自定义异常处理,如非法参数,则抛出异常
	if name == "hello" {
		//正确参数,return nil
		return nil
	} else {
		//非法参数,抛出自定义异常
		return errors.New("非法参数")
	}
}

//调用时,使用panic()输出异常,并终止程序
func test() {
	err := myHander("hell")
	if err != nil {
		//非法参数,内置函数panic()执行,输出异常,并终止程序
		panic(err)
	}
	//正常执行...
	fmt.Println("hello world")
}

func main() {
	test()
	fmt.Println("未出现异常,此代码执行...")
}
