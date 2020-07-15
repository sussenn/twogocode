package main

import "fmt"

type Point struct {
	x int
	y int
}
type Rect1 struct {
	leftUP, rightDown Point //Point对象作字段属性
}
type Rect2 struct {
	leftUP, rightDown *Point
}

func main() {
	r1 := Rect1{Point{1, 2}, Point{3, 4}}
	//r1是值类型, &r1取地址值
	fmt.Printf("r1.leftUP.x地址=%p\n r1.leftUP.y地址=%p\n r1.rightDown.x地址=%p\n r1.rightDown.y地址=%v\n",
		&r1.leftUP.x, &r1.leftUP.y, &r1.rightDown.x, &r1.rightDown.y)

	//r2是指针类型,传参&Point{}
	r2 := Rect2{&Point{10, 20}, &Point{30, 40}}

	//r2.leftUP地址=0xc0000381f0	 r2.rightDown地址=0xc0000381f8	--连续的
	fmt.Printf("r2.leftUP地址=%p\t r2.rightDown地址=%p\n",
		&r2.leftUP, &r2.rightDown)

	//r2.leftUP所指向的地址=0xc00000a0d0	 r2.rightDown所指向的地址=0xc00000a0e0	--不连续
	fmt.Printf("r2.leftUP所指向的地址=%p\t r2.rightDown所指向的地址=%p\n",
		r2.leftUP, r2.rightDown)

	//结论:
	//1. 结构体struct 字段在内存分布中是连续的
	//2. 指针类型字段其地址值在内存分布中是连续的, 但其所指向的地址值分布不是连续的
}
