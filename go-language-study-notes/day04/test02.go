package main

import "fmt"

func test1(a ...int) {
	for i := range a {
		a[i] += 100
	}
}
func main() {
	x := []int{1, 2, 3}
	test1(x...)    //注意这个地方是要加点的
	fmt.Println(x) //101 102 103
}
