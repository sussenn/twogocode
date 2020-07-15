package main

import "fmt"

func main() {
	//创建方式1-----------------------------
	var arr [4][6]int
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3
	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
	//创建方式2-----------------------------
	var arr3 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("arr3=", arr3)
	//遍历1
	for i := 0; i < len(arr3); i++ {
		for j := 0; j < len(arr3[i]); j++ {
			fmt.Printf("%v\t", arr3[i][j])
		}
		fmt.Println()
	}
	//遍历2
	for i, v := range arr3 {
		for j, v2 := range v {
			fmt.Printf("arr3[%v][%v]=%v\t", i, j, v2)
		}
		fmt.Println()
	}
}
