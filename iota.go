package main

import "fmt"

type flag byte

const (
	read  flag = 1 << iota //1
	write                  //2
	exec                   //4
)

func main() {
	fmt.Println(read)
	fmt.Println(write)
	fmt.Println(exec)
	//2的次方相加可以使用｜或运算简化
	fmt.Println(64 | 4)
}
