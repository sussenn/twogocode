package main

import "fmt"

//延时语句 延迟执行 该关键字修饰的语句
//1. defer 会将修饰的代码压入 栈 中, 先进后出.
//2. 入栈时会将 当前状态的变量一起拷贝进去! 即值不受后续代码影响
//3. 延时代码是在 当前函数结束前执行!
func sum(n1 int, n2 int) int {
	defer fmt.Println("第三打印", n1) //n1 n2 已入栈,所以不再受n++影响
	defer fmt.Println("第二打印", n2)
	n1++
	n2++
	res := n1 + n2
	fmt.Println("第一打印", res)
	return res
}

func main() {
	res2 := sum(1, 2)
	fmt.Println("第四打印", res2)
}
