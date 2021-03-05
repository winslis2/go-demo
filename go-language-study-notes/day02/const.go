package main

import "fmt"

func main() {
	x1 := new(int)
	fmt.Println(x1)
	//常量定义了可以不使用
	const x = 12

	const (
		a = 123
		//b默认和a是相同的
		b
	)
	fmt.Printf("%T,%v\n", b, b) //int,123

	//byte 和uint8 的类型是相同的，两者可以相互赋值
	var a1 uint8
	var a2 byte
	a1 = 2
	a2 = a1
	fmt.Println(a2)

}
