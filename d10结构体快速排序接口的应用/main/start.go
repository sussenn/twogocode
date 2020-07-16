package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//结构体切片排序 sort.Sort()方法的实现.	底层是快速排序
//1. sort.Sort()方法要求传参Interface接口,而Interface需要实现3个方法
//2. 使结构体切片HeroSlice 实现该接口3个方法,即可使用sort.Sort()方法
type Hero struct {
	Name string
	Age  int
}

//声明结构体切片
type HeroSlice []Hero

//实现排序接口方法
//1. Len()方法,返回List长度
func (hs HeroSlice) Len() int {
	return len(hs)
}

//2. Less()方法, 升序:ture
func (hs HeroSlice) Less(i, j int) bool {
	//Age升序,从小到大
	//return hs[i].Age < hs[j].Age
	//Name降序
	return hs[i].Name > hs[j].Name
}

//3. Swap()方法, 交换
func (hs HeroSlice) Swap(i, j int) {
	//temp := hs[i]
	//hs[i] = hs[j]
	//hs[j] = temp
	hs[i], hs[j] = hs[j], hs[i]
}

func main() {
	//声明结构体切片
	var heros HeroSlice
	//初始化 10个对象
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("%d-英雄", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		//将对象添加到集合
		heros = append(heros, hero)
	}
	//输出heros集合,即排序前的结果
	for _, v := range heros {
		fmt.Println(v)
	}
	fmt.Println("排序后--------------------")
	//传参Interface 接口, 而切片HeroSlice已实现其方法
	sort.Sort(heros)
	for _, vs := range heros {
		fmt.Println(vs)
	}
}
