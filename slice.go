package main

import "fmt"

func main() {
	//slice_test()
	slice_copy()
}
func slice_test() {
	//cap如果小len会直接报错
	s := make([]int, 2, 3) //超出容量也没事，会自动扩容
	s = append(s, 1)
	s = append(s, 2)
	nums := s[0:2]
	fmt.Println(s)
	fmt.Println(nums)
}

func slice_copy() {
	number := []int{1, 2, 3, 4}
	//var number1 []int
	//number1 :=make([]int,2) //这个只能复制前两个
	number1 := make([]int, 4)
	copy(number1, number)
	fmt.Println(number1)
}
