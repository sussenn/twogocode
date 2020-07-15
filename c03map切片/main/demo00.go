package main

import "fmt"

//map切片.	[]map
func main() {
	monsters := make([]map[string]string, 2)
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"
	}
	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "红孩儿"
		monsters[1]["age"] = "250"
	}
	//再添加一个元素,运行会报错: 数组越界
	//可以使用匿名函数append()动态追加!
	newMons := map[string]string{"name": "铁扇公主", "age": "400"}
	monsters = append(monsters, newMons)
	fmt.Println("切片map的创建,[]map=", monsters)
}
