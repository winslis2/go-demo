package main

import "fmt"

func main() {
	slice_test()
	//slice_copy()
}
func slice_test() {
	//cap如果小len会直接报错
	s := make([]int, 2, 3) //超出容量也没事，会自动扩容   这种情况前两个值是0，append 追加的是从index=2开始的
	//如果需要切面初始化	s:=int[]{1,2,3}
	s = append(s, 1)
	s = append(s, 2)
	for v := range s {
		fmt.Println(v)
	}
}

func slice_copy() {
	number := []int{1, 2, 3, 4}
	//var number1 []int
	//number1 :=make([]int,2) //这个只能复制前两个
	number1 := make([]int, 4)
	copy(number1, number)
	fmt.Println(number1)
}
