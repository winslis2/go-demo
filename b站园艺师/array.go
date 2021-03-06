package main

import "fmt"

func main() {
	//go是有数组的，slice是对数组的封装
	//数组作为参数效率是很低的，会进行复制，一般使用slice作为参数
	var planets [8]string
	planets[0] = "aaa"
	planets1 := []string{"aaa", "bbb"}
	fmt.Println(planets1[0])
}

func needArr([3]int) { //这个地方的参数是长度为3的数组

}
