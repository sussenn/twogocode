package main

import "fmt"

type Monster struct {
	Name string
	Age  int
}

//重写toString()
func (m *Monster) String() string {
	return fmt.Sprintf("{Name=%v,Age=%v}", m.Name, m.Age)
}

type E struct {
	Monster
	int //匿名字段
	num int
}

func main() {
	var e E
	e.Monster.Name = "tom"
	//此写法 等价于 e.Monster.Age
	e.Age = 18
	//匿名字段的访问
	e.int = 10
	e.num = 100
	//&e 只重写了Monster的String()所以输出Monster的值
	fmt.Println("匿名字段的访问:", e)
}
