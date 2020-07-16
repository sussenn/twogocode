package main

import "fmt"

//景区门票案例===============================================
type Visitor struct {
	Name string
	Age  int
}

func (visi *Visitor) showPrice() {
	if visi.Age >= 90 || visi.Age < 8 {
		fmt.Println("尊老爱幼,请离开")
		return
	}
	if visi.Age > 18 {
		fmt.Printf("游客年龄:%v\t 姓名:%v 门票20元\n", visi.Age, visi.Name)
	} else {
		fmt.Printf("游客年龄:%v\t 姓名:%v 免费\n", visi.Age, visi.Name)
	}
}

func main() {
	var v Visitor
	for {
		fmt.Println("请输入姓名:")
		//&v.Name --> &(v.Name)		"."的优先级更高
		fmt.Scanln(&v.Name)
		if v.Name == "n" {
			fmt.Println("退出程序...")
			break
		}
		fmt.Println("请输入年龄:")
		fmt.Scanln(&v.Age)
		v.showPrice()
	}
}
