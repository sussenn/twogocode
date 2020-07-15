package main

import (
	"fmt"
	operUtil "twogocode/a01fundemo/utils"
)

//可以给导入包起别名. 导包路径从src/后开始
func main() {
	var n1 = 1.2
	var n2 = 2.3
	var ope byte = '/'

	res := operUtil.Cal(n1, n2, ope)
	fmt.Println("结果值:", res)

}
