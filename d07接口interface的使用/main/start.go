package main

import "fmt"

//interface接口 默认是指针类型.		即未初始的接口,其默认值是 nil
//声明接口,只有方法	[不能包含任何变量]
type Usb interface {
	//声明方法
	Start()
	Stop()
}

//结构体Phone--------------------------------
type Phone struct {
}

//结构体Phone实现 Usb接口方法
func (p Phone) Start() {
	fmt.Println("手机 start...")
}
func (p Phone) Stop() {
	fmt.Println("手机 stop...")
}

//结构体TV----------------------------------
type TV struct {
}

//结构体TV实现 Usb接口方法
func (p TV) Start() {
	fmt.Println("电视 start...")
}
func (p TV) Stop() {
	fmt.Println("电视 stop...")
}

//结构体Computer-----------------------------
type Computer struct {
}

//结构体Computer的方法, 进行接口的调用
func (c Computer) Working(usb Usb) {
	//接收Usb接口作为入参,调用实现的方法
	usb.Start()
	usb.Stop()
}

func main() {
	//所有类型都实现了 空接口. 即任何变量都可以赋值给空接口
	//var xxx interface{}
	//------------------------------------------------------------
	com := Computer{}
	phone := Phone{}
	tv := TV{}
	//computer 不关心接口方法的具体实现,只需要调用即可.		至于方法的实现,针对具体行为另写方法即可
	com.Working(phone)
	com.Working(tv)
}
