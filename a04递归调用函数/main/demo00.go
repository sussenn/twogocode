package main

import "fmt"

//func test(n int) {
//func test2(n int) {
//	if n > 2 {
//		n--
//		//递归调用. 自己调自己
//		//test(n)
//		test2(n)
//	} else {
//		fmt.Println("n=", n)
//	}
//}
//
//func main() {
//	test2(4)
//}

//已知 fu(1)=3, fu(n)= 2*fu(n-1)+1
func fu(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*fu(n-1) + 1
	}
}
func main() {
	fmt.Println("fu(1)=", fu(1))
	fmt.Println("fu(5)=", fu(5))
}
