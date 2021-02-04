package main

import "fmt"

func main() {
	//pointer_test()
	pointer_test1()
}
func pointer_test() {
	var ptr *int
	fmt.Printf("%x", ptr) //0
	fmt.Println(ptr)      //nil

}

//指针数组  多个指针的地址，但是不存在这个数组指向一个地址，只有单独的指针才有地址
func pointer_test1() {
	num := []int{1, 2, 3}
	var ptr [3]*int
	for i := 0; i < len(num); i++ {
		ptr[i] = &num[i]
	}
	fmt.Println(*ptr[0], *ptr[1], *ptr[2]) //不能打印*ptr
}
