package main

import "fmt"

func main() {
	a := 1
	fmt.Printf("%p", &a) //0xc000016080  %p是指针，相比%x多了0x其他部分相同
	fmt.Printf("%x", &a) //c000016080
	//pointer_test()
	//pointer_test1()
}
func pointer_test() {
	var ptr *int
	fmt.Printf("%x", ptr) //0 指针ptr的值为0 %x16进制表示
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
