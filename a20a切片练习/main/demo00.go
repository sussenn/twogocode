package main

import "fmt"

func main() {
	//demo01------------------------------
	var sli1 = []int{10, 20, 30}
	sli2 := make([]int, 1)
	copy(sli2, sli1)
	//输出: [10]		虽然切片2长度只有1,但仍能进行copy.且仅能拷贝1个元素
	fmt.Println("长度为1的sli2,copy长度为3的sli1:", sli2)

}
