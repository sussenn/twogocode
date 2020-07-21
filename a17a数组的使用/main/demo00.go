package main

import "fmt"

//go中 数组是值类型, 是值拷贝
func main() {
	//定义一个4位长度的数组
	var hans [4]float64
	hans[0] = 1.0
	hans[1] = 1.0
	hans[2] = 2.0
	hans[3] = 3.0
	//遍历数组,并累加
	sum := 0.0
	for i := 0; i < len(hans); i++ {
		sum += hans[i]
	}
	//求平均值
	//%.2f :表示取float类型,小数点后2位
	//sum/? :如果是具体数字,即常量值,就不用强转类型.  而变量间的处理,必须统一类型
	avg := fmt.Sprintf("%.2f", sum/float64(len(hans)))
	fmt.Printf("总数=%v \t 平均值=%v \n", sum, avg)

	//==========================================
	//数组的初始化方式. 4种
	var arr01 [3]int = [3]int{1, 2, 3}
	var arr02 = [3]int{4, 5, 6}
	//[...]固定写法
	var arr03 = [...]int{7, 8, 9}
	//int{}内部指定坐标对应的值
	var arr04 = [...]int{1: 10, 0: 20, 2: 30}
	//类型推导
	strArr := [...]string{1: "苏", 0: "州", 2: "城"}

	fmt.Printf("arr01=%v\tarr02=%v\tarr03=%v\tarr04=%v\n", arr01, arr02, arr03, arr04)
	fmt.Println("strArr =", strArr)
	//===============================================
	//数组的 增强for循环
	names := [...]string{"宋江", "林冲", "李逵", "吴用"}
	for index, val := range names {
		fmt.Printf("数组坐标:%d\n", index)
		fmt.Printf("value=%v\n", val)
	}
}
