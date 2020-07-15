package main

import (
	"fmt"
	"sort"
)

//map是无序的.	没有专门方法进行排序,需要将map的k取出放入切片,然后排序
func main() {
	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90
	fmt.Println("无序map,map1=", map1)

	//定义切片 存放map的key
	var keys []int
	for k, _ := range map1 {
		keys = append(keys, k)
	}
	fmt.Println("无序map1的 k:", keys)
	//排序
	sort.Ints(keys)
	fmt.Println("排序后map1的 k:", keys)
	for _, k := range keys {
		fmt.Printf("map1[%v]=%v \n", k, map1[k])
	}
}
