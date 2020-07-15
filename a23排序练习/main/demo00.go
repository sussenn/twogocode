package main

import "fmt"

func main() {
	names := [5]string{"白眉鹰王", "金毛狮王", "紫衫龙王", "青翼蝠王"}
	var heroName = ""
	fmt.Println("请输入name:")
	fmt.Scanln(&heroName)

	//方式一:顺序查找
	for i := 0; i < len(names); i++ {
		if heroName == names[i] {
			fmt.Printf("找到:%v\t 下标:%d\n", names[i], i)
			break
		} else if i == (len(names) - 1) {
			fmt.Printf("找不到:%v\n", heroName)
		}
	}
	//方式二:
	index := -1
	for i := 0; i < len(names); i++ {
		if heroName == names[i] {
			index = i
			break
		}
	}
	if index != -1 {
		fmt.Printf("找到:%v\t 下标:%d\n", names[index], index)
	} else {
		fmt.Printf("找不到:%v\n", heroName)
	}
}
