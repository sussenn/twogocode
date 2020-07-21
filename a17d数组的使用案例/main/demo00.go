package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//1.添加字符A-Z,并输出==================
	var myChars [26]byte
	for i := 0; i < 26; i++ {
		myChars[i] = 'A' + byte(i)
	}
	for i := 0; i < 26; i++ {
		fmt.Printf("%c", myChars[i])
	}
	//2.取出数组内最大数,和其索引==================
	var intArr = [...]int{1, -1, 3, 80, 100}
	maxVal := intArr[0]
	maxValIndex := 0
	for i := 0; i < len(intArr); i++ {
		if maxVal < intArr[i] {
			maxVal = intArr[i]
			maxValIndex = i
		}
	}
	fmt.Printf("\n最大数=%v \t索引=%v\n", maxVal, maxValIndex)
	//3.求数组整数和,以及平均数==========================
	var intArr2 = [...]int{1, -1, 3, 80, 10}
	sum := 0
	for _, val := range intArr2 {
		sum += val
	}
	fmt.Printf("最大数=%v\t 平均数=%v\n", sum, sum/len(intArr2))
	//4.随机5位数,并将其位置反转======================================
	var intArr3 [5]int
	len := len(intArr3)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		intArr3[i] = rand.Intn(100) //随机数范围:[0 100)
	}
	fmt.Println("生成的随机数=", intArr3)
	//开始数组数值位置的反转交换
	temp := 0
	//交换次数是 len/2
	for i := 0; i < len/2; i++ {
		//将最后一位(6-1-0=坐标5)与0位交换
		temp = intArr3[len-1-i]
		intArr3[i] = temp
	}
	fmt.Println("反转后的随机数=", intArr3)
}
