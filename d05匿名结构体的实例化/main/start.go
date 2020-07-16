package main

import "fmt"

type Goods struct {
	Name  string
	Price float64
}
type Brand struct {
	Name    string
	Address string
}

//这是多重继承!
type TV struct {
	Goods
	Brand
}
type TV2 struct {
	*Goods
	*Brand
}

func main() {
	tvone := TV{Goods{"显示器", 1000.00}, Brand{"创维", "杭州"}}
	fmt.Println("匿名结构体的值传递实例化:", tvone)

	//Goods和Brand 是指针类型
	//tvtwo的2个参数 指向 Goods和Brand的地址值--> {0xc0000044c0 0xc0000044e0}
	tvtwo := TV2{&Goods{Name: "空调", Price: 500.00}, &Brand{Name: "格力", Address: "北京"}}
	//tvtwo本身不是指针类型!	所以以下写法错误!
	//fmt.Println(*tvtwo)	//err
	fmt.Println("匿名结构体的地址值传递实例化:", *tvtwo.Goods, *tvtwo.Brand)

}
