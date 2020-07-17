package main

import "fmt"

type Monkey struct {
	Name string
}

//2个接口
type BirdAble interface {
	Flying()
}
type FishAble interface {
	Swimming()
}

func (monkey *Monkey) climbing() {
	fmt.Println(monkey.Name, "继承父类,天生会爬树...")
}

type LittleMonkey struct {
	//继承
	Monkey
}

func (litMonkey *LittleMonkey) Flying() {
	fmt.Println(litMonkey.Name, "通过实现接口方法,学会飞...")
}

func (litMonkey *LittleMonkey) Swimming() {
	fmt.Println(litMonkey.Name, "通过实现接口方法,学会游泳...")
}

func main() {
	monkey := LittleMonkey{
		Monkey{
			Name: "孙悟空",
		},
	}
	monkey.climbing() //继承父类的方法
	monkey.Flying()
	monkey.Swimming()
}
