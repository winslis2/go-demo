package main

import "fmt"

func main() {
	//slice是指向数组的窗口/视图
	//slice 可以切字符串[1:4]  [:]

	drarfs := []string{"1", "c"} //声明slice 有两个方式，一个是slice一个数组，二是这种不用给数组加长度
	fmt.Println(drarfs)
	//动态扩展后会变成两个数组
	a := make([]int, 1, 1) //默认是[0]
	ints1 := append(a, 1)  //[0,1]
	ints2 := append(a, 2)  //[]
	fmt.Println(ints1)
	ints2[0] = 11
	fmt.Println(ints2)
	fmt.Println(ints1)
	//append可一次增加多个

}

//第三个表示容量
func threeSlice() {
	//planets :=[]string{"1","2","3","4","5","6","7","8"}
	////planets2修改后planet的第5个会发生变化
	//planets1 := planets[0:4]
	//planets2 := append(planets1, "6")
	//
	////planets2修改后planet的第5个不会发生变化
	//planets1 := planets[0:4:4]
	//planets2 := append(planets1, "6")
}
