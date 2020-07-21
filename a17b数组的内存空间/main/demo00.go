package main

import "fmt"

//数组是	值类型
func main() {
	var intArr [3]int //int64 占位8个字节,	int32占位4个字节
	fmt.Println("数组的初始值:", intArr)
	intArr[0] = 10
	intArr[1] = 20
	intArr[2] = 30
	fmt.Println("数组赋值后:", intArr)
	//1.数组的地址值=数组首位地址值. 即intArr地址值=intArr[0]地址值
	//2.数组地址值,相邻2位相差一个类型字节占位空间. 如intArr[0]地址值=0xc000010380	|	intArr[1]地址值=0xc000010388
	//3.地址值是16进制,所以intArr[1]地址值=0xc000010388		intArr[2]地址值=0xc000010390
	fmt.Printf("intArr地址值=%p \tintArr[0]地址值=%p \tintArr[1]地址值=%p \tintArr[2]地址值=%p \n",
		&intArr, &intArr[0], &intArr[1], &intArr[2])
}
