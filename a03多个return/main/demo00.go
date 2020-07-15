package main

import "fmt"

//2个返回值
func getSumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func main() {
	res1, res2 := getSumAndSub(1, 2)
	fmt.Printf("res1 = %v, res2 = %v \n", res1, res2)

	_, res3 := getSumAndSub(5, 6)
	fmt.Println("res3 = ", res3)
}
