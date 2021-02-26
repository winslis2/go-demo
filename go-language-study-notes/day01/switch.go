package main

import "fmt"

func main() {
	x := 100
	//swtich后可以什么都不跟
	switch {
	case x < 100:
		fmt.Println("x<100")
	case x >= 100:
		fmt.Println("x>=100")
	}
	//也可以不写default
}
