package main

import "fmt"

//slice是引用类型
func main() {
	//先定义一个数组
	var intArr [5]int = [...]int{1, 22, 33, 55, 99}
	//声明切片mySli	取intArr 1-3位[1 3)
	mySli := intArr[1:3]
	//[22,33]
	fmt.Println("切片元素是:", mySli)
	//cap=4		容量:切片当前能存储元素的最大值. 其是动态变化的
	fmt.Println("切片容量是:", cap(mySli))
	//获取intArr[1]和mySli[0]地址值
	//intArr[1]=0xc00000c3f8	 mySli[0]=0xc00000c3f8
	fmt.Printf("地址值:intArr[1]=%p\t mySli[0]=%p\n", &intArr[1], &mySli[0])
	//mySli=0xc0000044a0
	//切片与数组的差异: &arr=&arr[0]	毕竟&slice[0]只是slice一部分,指向&arr[0]即维护了一个数组,进行数据存储
	fmt.Printf("地址值:mySli=%p\n", &mySli)

	//结论:
	//1.切片slice是引用类型
	//2.切片slice的内存布局 包含3部分:指针ptr(对应数组首位元素intArr[1]的地址值)/切片长度len/切片容量cap
	//3.切片slice是一个struct结构体
}
