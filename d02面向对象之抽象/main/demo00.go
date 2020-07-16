package main

import "fmt"

type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64 //余额
}

//存款
func (account *Account) Deposite(money float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("密码输入有误")
		return
	}
	if money <= 0 {
		fmt.Println("请输入正确金额")
		return
	}
	account.Balance += money
	fmt.Printf("存款成功,存款金额:%v\n", money)
}

//取款
func (account *Account) WithDraw(money float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("密码输入有误")
		return
	}
	if money <= 0 || money > account.Balance {
		fmt.Println("余额不足")
		return
	}
	account.Balance -= money
	fmt.Printf("取款成功,取款金额:%v\n", money)
}

//查询余额
func (account *Account) Query(pwd string) {
	if pwd != account.Pwd {
		fmt.Println("密码输入有误")
		return
	}
	fmt.Printf("账户:%v\t 余额:%v\n", account.AccountNo, account.Balance)
}

func main() {
	account := Account{"No001", "123", 100.00}
	account.Query("123")
	account.Deposite(100.00, "123")
	account.WithDraw(50.00, "123")
	account.Query("123")
}
