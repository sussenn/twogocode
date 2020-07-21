package main

import "fmt"

//slice的创建
func main() {
	//1.定义切片,然后去引用已创建的数组.		具体参考b09...
	//-----------------------------------------------------
	//2.使用内置函数make()创建
	//语法: make([]类型, len, cap)		//cap可选. 如定义,则要求cap>=len
	var slice []float64 = make([]float64, 5, 10)
	slice[1] = 10
	slice[3] = 20
	fmt.Println(slice)
	fmt.Println("slice的长度:", len(slice)) //5
	fmt.Println("slice的容量:", cap(slice)) //10
	//使用make()创建的切片slice, 内存布局和方式一是一致的.
	//即slice引用数组arr[5],slice[0]地址值指向arr[0]	(底层维护arr数组对外不可见)
	//本身包括3部分: slice[0]地址值/len/cap	详见b09...
	//----------------------------------------------------
	//3.
	var strSlice []string = []string{"cat", "tom", "andy"}
	fmt.Println(strSlice)
	fmt.Println("strSlice的长度:", len(strSlice)) //3
	fmt.Println("strSlice的容量:", cap(strSlice)) //3
	//======================================================
	//针对方式一,可以简写
	var arr = [...]int{0, 1, 2, 3, 4}
	//mySli := arr[:3]		//取0-3位
	//mySli := arr[1:]		//取1-len位
	mySli1 := arr[:] //取0-len位
	fmt.Println(mySli1)
	//------------------------------------------------
	//针对原有切片可进行再次切片
	//mySli2[0]指向arr[1]地址值
	mySli2 := mySli1[1:3]
	fmt.Println(mySli2)
	//-------------------------------------------------
	//使用内置函数append()对原有切片进行元素追加
	//底层原理: 使用append()会创建新的arr[],将切片原先维护的数组内容拷贝到新的arr[],然后将切片重新指向新数组arr[]地址
	mySli3 := []int{0, 1, 2, 3, 4, 5}
	mySli3 = append(mySli3, 6, 7)
	fmt.Println(mySli3)
	//二次追加	"..."固定写法
	mySli3 = append(mySli3, mySli3...)
	fmt.Println(mySli3)
	//-------------------------------------------------
	//内置函数copy()
	var mySli4 = []int{0, 1, 2, 3, 4}
	mySli5 := make([]int, 10)
	//将mySli4复制到mySli5.		copy()入参必须是2个切片类型
	copy(mySli5, mySli4)
	//copy后,两个切片空间是相互独立的. 即mySli4重新赋值,不影响mySli5
	mySli4[0] = 100
	fmt.Println("copy后的mysli4:", mySli4)
	fmt.Println("copy后的mysli5:", mySli5)
}
