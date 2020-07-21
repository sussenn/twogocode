package main

import "fmt"

func main() {
	//普通for循环
	var arr = [...]int{1, 2, 3, 4, 5}
	slice := arr[1:4]
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]=%v\n", i, slice[i])
	}
	//增强for循环
	for i, val := range slice {
		fmt.Printf("slice[%v]=%v\n", i, val)
	}
	//==================================================

}
