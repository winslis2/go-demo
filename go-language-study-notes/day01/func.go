package main

import (
	"errors"
	"fmt"
)

func main() {
	x, _ := div(10, 2)
	fmt.Println(x)
}

//go 可以返回多个值，类型定义是放到后面的
//java和PHP只能有一个返回值，数据类型是放到前面的
func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

//func() 可作为返回类型 同时这个返回函数类型也可以规定返回类型，和内部的函数类型是相同的
func ret_fun() func() string {
	return func() string {
		return "ss"
	}
}
