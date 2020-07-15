package main

import "fmt"

//冒泡排序
func BubbleSort(arr *[5]int) [5]int {
	fmt.Println("排序前arr=", *arr)
	temp := 0
	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}
	return *arr
}
func main() {
	arr := [5]int{38, 26, 99, 41, 19}
	//传入指针类型(地址值), 返回值类型
	arr = BubbleSort(&arr)
	fmt.Println("排序后arr=", arr)
}
