package main

import (
	"fmt"
	"strconv"
)

func main() {
	//本来有个疑问，为啥使用包里的函数的总是大写的，原来是因为大写的是public
	//字符串转数字
	a, _ := strconv.ParseInt("11", 2, 64)
	b, _ := strconv.ParseInt("0144", 8, 64)
	c, _ := strconv.ParseInt("64", 16, 64)
	fmt.Println(a, b, c)

	//数字转化成字符串
	fmt.Println(strconv.FormatInt(a, 2))
	fmt.Println(strconv.FormatInt(b, 8))
	fmt.Println(strconv.FormatInt(c, 16))
}
