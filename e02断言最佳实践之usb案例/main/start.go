package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}
type Phone struct {
	Name string
}

func (p Phone) Start() {
	fmt.Println(p.Name, "手机开始工作...")
}
func (p Phone) Stop() {
	fmt.Println(p.Name, "手机停止工作...")
}

//Phone自己的方法
func (p Phone) Call() {
	fmt.Println(p.Name, "在打电话.")
}

type Camera struct {
	Name string
}

func (c Camera) Start() {
	fmt.Println(c.Name, "相机开始工作...")
}
func (c Camera) Stop() {
	fmt.Println(c.Name, "相机停止工作...")
}

type Computer struct {
}

func (computer Computer) Working(usb Usb) {
	usb.Start()
	//如果传参是"Phone"类型,则会调用Phone独有方法
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop()
}

func main() {
	var usbArr [3]Usb
	usbArr[0] = Phone{"华为"}
	usbArr[1] = Phone{"苹果"}
	usbArr[2] = Camera{"尼康"}

	var computer Computer
	for _, v := range usbArr {
		computer.Working(v)
	}
}
