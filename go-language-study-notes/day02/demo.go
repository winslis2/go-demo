package main

import "fmt"

func main() {
	a := 2
	p := &a
	//指针类型不能使用int强转   指针类型=*int
	fmt.Println((*int)(p)) //为了避免歧义，两个都要括起来
	fmt.Println(&a)
}
