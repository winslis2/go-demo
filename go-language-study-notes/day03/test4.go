package main

func main() {
	a, b, c, x := 1, 2, 3, 2
	switch x {
	case a, b: //x和a，或b比较
		println("a|b")
	case c:
		println("c")
	default:
		println("d")
	}
}
