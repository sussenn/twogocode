package main

import "fmt"

//如果某个用户名存在，就将其密码修改"888888"，如果不存在就增加这个用户信息,（包括昵称nickname 和 密码pwd）
func modifyUser(users map[string]map[string]string, name string) {
	if users[name] != nil {
		//有这个用户
		users[name]["pwd"] = "888888"
	} else {
		//没有这个用户
		users[name] = make(map[string]string, 2)
		users[name]["pwd"] = "888888"
		users[name]["nickname"] = "昵称~" + name //示意
	}
}
func main() {
	users := make(map[string]map[string]string, 10)
	users["smith"] = make(map[string]string, 2)
	users["smith"]["pwd"] = "999999"
	users["smith"]["nickname"] = "小花猫"
	modifyUser(users, "tom")
	modifyUser(users, "mary")
	modifyUser(users, "smith")
	fmt.Println(users)
}
