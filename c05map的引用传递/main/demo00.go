package main

import "fmt"

func modify(map1 map[int]int) map[int]int {
	map1[10] = 999
	return map1
}

func main() {
	map1 := make(map[int]int)
	map1[1] = 60
	map1[2] = 30
	map1[10] = 1
	map1[20] = 2
	//值改变		map是引用类型
	map1 = modify(map1)
	fmt.Println(map1)
}
