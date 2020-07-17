package main

import "fmt"

func TypeJudge(items ...interface{}) {
	for index, v := range items {
		switch v.(type) {
		case bool:
			fmt.Printf("第%v参数是bool类型, 值:%v\n", index, v)
		case float64:
			fmt.Printf("第%v参数是float64类型, 值:%v\n", index, v)
		case int, int32, int64:
			fmt.Printf("第%v参数是整数类型, 值:%v\n", index, v)
		case string:
			fmt.Printf("第%v参数是string类型, 值:%v\n", index, v)
		default:
			fmt.Printf("第%v参数类型不确定, 值:%v\n", index, v)
		}
	}
}

func main() {
	var n1 int32 = 18
	var n2 float64 = 31.8
	var name = "tom"
	address := "深圳"
	n3 := 100
	TypeJudge(n1, n2, name, address, n3)
}
