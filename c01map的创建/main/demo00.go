package main

import "fmt"

//map是引用类型.	map的创建	k-v无序
func main() {
	//方式一--------------------------------
	//map的声明
	var names map[string]string
	//map的初始化,即分配内存空间
	names = make(map[string]string, 10)
	names["no1"] = "令狐冲"
	names["no2"] = "叶孤城"
	names["no1"] = "西门吹雪"
	names["no3"] = "东方不败"
	fmt.Println("map创建方式一.map=:", names)
	//方式二------------------------------
	cities := make(map[string]string)
	cities["01"] = "北京"
	cities["02"] = "上海"
	cities["03"] = "深圳"
	fmt.Println("map创建方式二.map=:", cities)
	//方式三-------------------------------
	man := map[string]string{"m1": "孙悟空", "m2": "猪八戒", "m3": "沙和尚"}
	man["m4"] = "唐僧"
	fmt.Println("map创建方式三.map=:", man)
	//======================================================================
	//map作value
	stuMap := make(map[int]map[string]string)

	stuMap[1] = make(map[string]string, 3)
	stuMap[1]["name"] = "tom"
	stuMap[1]["sex"] = "男"
	stuMap[1]["address"] = "上海"

	stuMap[2] = make(map[string]string, 3) //此处必须再次初始化内存空间!
	stuMap[2]["name"] = "anny"
	stuMap[2]["sex"] = "女"
	stuMap[2]["address"] = "深圳"

	fmt.Println("stuMap=", stuMap)
	//map[address:深圳 name:anny sex:女]
	fmt.Println("stuMap[2]=", stuMap[2])
	//深圳
	fmt.Println("stuMap[2]的address=", stuMap[2]["address"])
	//========================================================================
	fmt.Println("-------------------------------------")
	//struct结构体作value
	studenrs := make(map[string]Stu, 10)
	stu1 := Stu{"tom", 18, "北京"}
	stu2 := Stu{"suu", 38, "东京"}
	studenrs["no1"] = stu1
	studenrs["no2"] = stu2

	for k, v := range studenrs {
		fmt.Printf("学生的编号是:%v \n", k)
		fmt.Printf("学生的名字是:%v \n", v.Name)
		fmt.Printf("学生的年龄是:%v \n", v.Age)
		fmt.Printf("学生的地址是:%v \n", v.Addr)
		fmt.Println()
	}

}

//定义结构体
type Stu struct {
	Name string
	Age  int
	Addr string
}
