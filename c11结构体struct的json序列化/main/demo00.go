package main

import (
	"encoding/json"
	"fmt"
)

type Man struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	m1 := Man{"奥特曼", 888, "动感光波"}
	//json序列化
	m1Str, err := json.Marshal(m1)
	if err != nil {
		fmt.Println("json解析错误:", err)
	} else {
		fmt.Println("json序列化后:", m1Str)
	}
}
