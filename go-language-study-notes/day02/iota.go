package main

import "fmt"

const a2 = 10

func main() {
	//fmt.Println(&a2)// can't take the address of 'a2'  数字常量不会分配内存地址
	const (
		a = iota
		b
		c
	)
	//iota是从0开始的
	fmt.Println(a, b, c)

	//可以使用空白符先用一个
	const (
		_ = iota
		bb
		cc
	)
	fmt.Println(bb, cc)

	const (
		a1 = iota
		b1
		c1 = 100
		d1 = iota //上面的b1 c1无论怎么写都会造成iota自增
		e1
	)
	fmt.Println(a1, b1, c1, d1, e1) //0 1 100 3 4
}
