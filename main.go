package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//var s string = "hello world"
	//var b bool = false
	//var ch byte = 'a'
	//自动类型推断，放在函数中使用
	//test1, test2 := false,"hell world"
	//fmt.Println(test1)
	//fmt.Println(test2)
	//取地址
	//var pointer string = "pointer"
	//fmt.Println(&pointer)
	//空白符_的使用
	//_, i2, s := numbers()
	//fmt.Println(i2,s)
	//const的使用
	//const_test()
	//指针
	//num := 123
	//var ptr *int
	//ptr = &num
	//fmt.Println(*ptr)
	//if
	//if_test()
	//switch
	//switch_test()
	//fallthrough
	//fallthrough_test()
	//for
	//for_test()
	//for_range
	//for_range_test()
	//swap
	//a := 1
	//b := 2
	//swap(&a, &b)
	//fmt.Print(a, b)
	//闭包
	//sequence := getSequence() //返回的是个函数
	//fmt.Println(sequence())
	//fmt.Println(sequence())
	//数组
	arr_test()
}

//空白符测试函数
func numbers() (int, int, string) {
	a, b, c := 1, 2, "hello"
	return a, b, c
}

func const_test() {
	const COUNTY = "China"
	//const 是没有：的，类型可以省去
	//多个const写法
	const (
		a = "abc"
		b = len(a)
		c = unsafe.Sizeof(a)
	)
	fmt.Println(COUNTY)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

// if 语句执行
func if_test() {
	num := 10
	if num == 10 {
		fmt.Println("num等于10")
	} else {
		fmt.Println("num不等于10")
	}
}

// switch 语句执行
func switch_test() {
	age := 20
	switch age {
	case 20:
		fmt.Println("年轻人")
	default:
		fmt.Println("非年轻人")
	}
}

func fallthrough_test() {
	//fallthrough会强制执行下一个case,不管true或false

	switch {
	case false:
		fmt.Println("1、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	case false:
		fmt.Println("3、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("4、case 条件语句为 true")
	//这个地方没有写fallthrough所以不会执行后面的句子
	case false:
		fmt.Println("5、case 条件语句为 false")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}
}

func for_test() {
	//go是没有while循环的额
	sum := 0 //下面这种for写法要从1开始
	//for i:=0;i<100;i++ {
	//	sum = sum+i
	//}
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}
}

func for_range_test() {
	strings := []string{"string1", "string2"}
	nums := [5]int{1, 2, 4}
	for i, j := range strings {
		fmt.Println(i, j)
	}
	for i, j := range nums {
		fmt.Println(i, j)
	}
}

func swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

//因为返回值类型是func 且这个返回函数的返回类型是int
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

//返会类型对比
func func_type() func() {
	return func() {
	}
}

func arr_test() {
	//没有限定范围可以使用append动态添加
	numbs := []int{} //实际上这是切片
	numbs = append(numbs, 1)
	numbs = append(numbs, 2)
	fmt.Println(numbs)
	fmt.Println(len(numbs))
}
