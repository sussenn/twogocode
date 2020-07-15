package main

import "fmt"

//...int 可变参数,可输入多个值
func sum(n1 int, args ...int) int {
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func main() {
	res := sum(10, 0, 1, 2, 3)
	fmt.Println("res=", res)
}
