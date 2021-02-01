package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//var s string = "hello world"
	//var b bool = false
	//var ch byte = 'a'
	//自动类型推断，放在函数中使用
	//test1, test2 := false,"hell world"
	//fmt.Println(test1)
	//fmt.Println(test2)
	//取地址
	//var pointer string = "pointer"
	//fmt.Println(&pointer)
	//空白符_的使用
	//_, i2, s := numbers()
	//fmt.Println(i2,s)
	//const的使用
	//const_test()
	//指针
	num := 123
	var ptr *int
	ptr = &num
	fmt.Println(*ptr)
}

//空白符测试函数
func numbers() (int, int, string) {
	a, b, c := 1, 2, "hello"
	return a, b, c
}

func const_test() {
	const COUNTY = "China"
	//const 是没有：的，类型可以省去
	//多个const写法
	const (
		a = "abc"
		b = len(a)
		c = unsafe.Sizeof(a)
	)
	fmt.Println(COUNTY)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
