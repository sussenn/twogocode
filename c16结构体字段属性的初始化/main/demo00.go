package main

import "fmt"

//结构体内部属性的初始化=======
type Stu struct {
	Name string
	Age  int
}

//实现String()
func (s *Stu) String() string {
	str := fmt.Sprintf("Stu{Name=%v,Age=%v}", s.Name, s.Age)
	return str
}

func main() {
	var stu1 = Stu{"tom", 18}
	stu2 := Stu{"mark", 28}
	var stu3 = Stu{Name: "any", Age: 38}
	stu4 := Stu{Name: "jack", Age: 48}
	//&stu4: Stu{Name=jack,Age=48}
	fmt.Println(stu1, stu2, stu3, &stu4)

	//获取指针类型的结构体,并直接初始化-----------------------------------------
	var stu5 *Stu = &Stu{"aaa", 10}
	stu6 := &Stu{"bbb", 20}
	var stu7 = &Stu{"ccc", 30}
	stu8 := &Stu{"ddd", 40}
	//[注意:] 如果没有实现String(),则需要"*"获取指向值!
	fmt.Println(stu5, stu6, stu7, stu8)
}
