package main

import (
	"fmt"
	"strconv"
)

func main() {
	a, _ := strconv.ParseInt("1100100", 2, 64)
	fmt.Printf("%T\n", a)
	formatInt := strconv.FormatInt(a, 2) //转化成2进制字符串
	fmt.Println(formatInt)

}
