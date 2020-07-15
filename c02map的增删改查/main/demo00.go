package main

import "fmt"

func main() {
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "深圳"
	//map的查找----------------------------------------
	val, ok := cities["no4"]
	if ok {
		fmt.Println("根据key查找value,val=", val)
	} else {
		//找不到则无返回
		fmt.Println("根据key查找value,找不到则无返回,val=", val)
	}

	//内置函数-根据key删除-------------------------------
	delete(cities, "no1")
	//如key不存在,则不操作
	delete(cities, "no4")
	fmt.Println("map删除操作演示,map=", cities)

	//全部删除的方式: 1.遍历删除 2.将map置空
	cities = make(map[string]string) //map置空,无引用则被gc回收
	fmt.Println("将map置空,map=", cities)
	fmt.Println("------------------------")

	//==========================================================
	//map的遍历
	address := make(map[string]string)
	address["no1"] = "江西"
	address["no2"] = "安徽"
	address["no3"] = "浙江"
	for k, v := range address {
		fmt.Printf("遍历map,k = %v, v = %v\n", k, v)
	}
	//--------------------------------------------
	//复杂map的遍历
	stuMap := make(map[int]map[string]string)

	stuMap[1] = make(map[string]string, 3)
	stuMap[1]["name"] = "tom"
	stuMap[1]["sex"] = "男"
	stuMap[1]["address"] = "上海"

	stuMap[2] = make(map[string]string, 3) //此处必须再次初始化内存空间!
	stuMap[2]["name"] = "anny"
	stuMap[2]["sex"] = "女"
	stuMap[2]["address"] = "深圳"

	for k1, v1 := range stuMap {
		fmt.Println("复杂map的遍历,k1=", k1)
		fmt.Println("复杂map的遍历,v1=", v1)
		for k2, v2 := range v1 {
			fmt.Printf("复杂map的遍历,k2=%v,v2=%v\n", k2, v2)
		}
	}
}
