package main

import "fmt"

func main() {
	//map_test()
	map_test1()
}
func map_test() {
	man := make(map[string]string)
	man["name"] = "lis2"
	man["city"] = "henan"
	man["face"] = "good"
	//遍历出的结果顺序是不同，如果只写一个默认是key  slice同理，如果要的到value不想得到key就要使用_代替
	//for key ,value:=range man{
	//	fmt.Println(key,value)
	//}
	s, ok := man["ii"] // 直接输出不会报错是没有值的
	if ok {
		fmt.Println("存在")
	} else {
		fmt.Println("不存在")
	}
	fmt.Println(s)
}

func map_test1() {
	//初始化map map后不是跟的括号
	man := map[string]string{"name": "lis2", "city": "henan"}
	fmt.Println(man)
}
