package main

import "fmt"

//func main() {
//	res := func(n1 int, n2 int) int {
//		return n1 + n2
//	}(10, 20)
//	fmt.Println(res)
//}

//func main() {
//	noname := func(n1 int, n2 int) int {
//		return n1 + n2
//	}
//	res2 := noname(10, 20)
//	fmt.Println(res2)
//}

var (
	global_noname1 = func(n1 int, n2 int) int {
		return n1 + n2
	}
)

func main() {
	res3 := global_noname1(10, 10)
	fmt.Println(res3)
}
