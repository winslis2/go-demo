package main

import "fmt"

func main() {
	//会先执行case
	switch x := 5; x {
	default:
		x += 100
		fmt.Println(x)
	case 5:
		x += 50
		fmt.Println(x)
	}
	//fmt.Println(x)读不到x的值
}
