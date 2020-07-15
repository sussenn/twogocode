package main

import (
	"fmt"
	"strconv" //字符串和整数的处理
	"strings"
)

func main() {
	str1 := "hel 杭州"
	//字符串的遍历	将含中文字符串 转 []rune
	res1 := []rune(str1)
	for i := 0; i < len(res1); i++ {
		fmt.Printf("%c\n", res1[i])
	}
	//字符串转整数---------------------------
	n, err := strconv.Atoi("123")
	//n, err := strconv.Atoi("hello")
	//nil -> 类似null	如不处理 则打印输出 <nil> 占位
	if err != nil {
		fmt.Println("转换错误:", err)
	} else {
		fmt.Println("字符串转换整数:", n)
	}
	//整数转字符串---------------------------
	str2 := strconv.Itoa(321)
	fmt.Println("整数转换字符串:", str2)
	//字符串转字节数组------------------------
	var myByte = []byte("hel深圳")
	fmt.Println("字符串转换字节数组:", myByte)
	//字节数组转字符串------------------------
	str3 := string([]byte{97, 98, 99})
	fmt.Println("字节数组转换字符串:", str3)
	//10进制转2,8,16..进制-------------------
	str4 := strconv.FormatInt(123, 2)
	fmt.Printf("十进制转二进制:%v \n", str4)
	str5 := strconv.FormatInt(123, 8)
	fmt.Printf("十进制转八进制:%v \n", str5)
	//==========================================================================
	//判断指定字符是否存在----------------------
	boo1 := strings.Contains("hello[]world", "[]")
	fmt.Println("判断指定字符是否存在:", boo1)
	//统计指定字符有多少个----------------------
	num := strings.Count("hellooo", "o")
	fmt.Println("统计指定字符有多少个:", num)
	//字符串内容的比较-------------------------
	boo2 := strings.EqualFold("abc", "ABC")
	fmt.Println("字符串内容的比较(不区分大小写):", boo2)
	fmt.Println("字符串内容的比较(区分大小写):", "abc" == "ABC")
	//返回指定字符串首次出现的地标(-1表示不存在)---------------
	index1 := strings.Index("HELLO_abcdef", "abc")
	fmt.Println("指定字符串首次出现的位数", index1)
	//返回指定字符串最末次出现的地标(-1表示不存在)---------------
	index2 := strings.LastIndex("HELLO_abcdefabc", "abc")
	fmt.Println("指定字符串最末次出现的位数", index2)
	//指定字符串的替换(-1表示替换所有)--------------------------------------
	str6 := strings.Replace("helloworldo", "o", "A", -1)
	fmt.Println("替换指定的字符串", str6)
	//字符串的切割,切割后变成字符串数组-----------------------------------
	strArr1 := strings.Split("hello,world,周", ",")
	fmt.Println("字符串的切割后", strArr1)
	for i := 0; i < len(strArr1); i++ {
		fmt.Printf("strArr[%d] = %v \t", i, strArr1[i])
	}
	//字符串的大小写转换-----------------------------------------------
	str7 := strings.ToLower("hello WORLD")
	fmt.Printf("\n字符串的大写转小写: %v \n", str7)
	str8 := strings.ToUpper("hello WORLD")
	fmt.Println("字符串的小写转大写:", str8)
	//去掉字符串两边空格-----------------------------------------------
	str9 := strings.TrimSpace("  nihao hangzhou ")
	fmt.Println("去掉字符串两边空格:", str9)
	//去掉字符串两边指定字符及空格---------------------------------------
	str10 := strings.Trim("@ gmail.com@", " @")
	fmt.Println("去掉字符串两边指定字符及空格:", str10)
	//判断字符串是否以指定字符串开头/结尾-------------------------------------
	boo3 := strings.HasPrefix("https://localhost:8080", "https://")
	fmt.Println("判断字符串是否以指定字符串开头:", boo3)
	boo4 := strings.HasSuffix("资料.xslx", ".xslx")
	fmt.Println("判断字符串是否以指定字符串结尾:", boo4)
}
