package main

import (
	"fmt"
	"time"
)

//go 内置常用的时间常量. 无法使用 秒/1000 获取毫秒!
func main() {
	now := time.Now()

	time.Sleep(time.Nanosecond * 100)  //休眠100纳妙
	time.Sleep(time.Microsecond * 100) //休眠100微妙
	time.Sleep(time.Millisecond * 100) //休眠100毫秒
	time.Sleep(time.Second * 100)      //休眠100秒
	time.Sleep(time.Minute * 100)      //休眠100分

	//时间戳'方法'	获取1970-01-01至今的毫秒/纳秒数
	fmt.Println("unix毫秒时间戳:", now.Unix())
	fmt.Println("unix纳秒时间戳:", now.UnixNano())
}
