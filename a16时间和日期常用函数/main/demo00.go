package main

import (
	"fmt"
	"time"
)

func main() {
	//获取当前时间 	类型: time.Time
	now := time.Now()
	fmt.Printf("当前时间: %v, 类型: %T \n", now, now)
	fmt.Printf("年=%v\t月=%v\t日=%v\t时=%v\t分=%v\t秒=%v\n", now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
	//日期格式化一
	dateStr := fmt.Sprintf("当前时间格式化: %d-%d-%d %d:%d:%d", now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println("dateStr=", dateStr)
	//日期格式化二	"2006-01-02 15:04:05" 固定写法!相当于yyyy-MM-dd HH:mm:ss	[推荐]
	fmt.Println(now.Format("2006-01-02 15:04:05"))
}
