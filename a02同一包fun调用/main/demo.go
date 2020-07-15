package main

import "fmt"

func test(n1 int) {
	n1 = n1 + 1
	fmt.Println("test() n1 = ", n1)
}

func getSum(n1 int, n2 int) int {
	sum := n1 + n2
	fmt.Println("getSum() sum = ", sum)
	return sum
}

func main() {
	n1 := 10
	test(n1)
	fmt.Println("main() n1 = ", n1)

	sum := getSum(10, 20)
	fmt.Println("main() sum = ", sum)

}
