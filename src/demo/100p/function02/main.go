package main

import "fmt"

//func test(n1 int) {
//	n1 = n1 + 1
//	fmt.Println("n1=", n1)
//}

//func main() {
//	n1 := 10
//	test(n1)
//	fmt.Println("n1=", n1)
//}
//
//func getSum(n1 int, n2 int) int {
//	sum := n1 + n2
//	return sum
//}
//
//func main() {
//	sum := getSum(10, 20)
//	fmt.Println("sum=", sum)
//}

func test(n1 float64, n2 float64) (float64, float64) {
	he := n1 + n2
	cha := n1 - n2
	return he, cha
}

func main() {
	n1, n2 := 10.88, 1024.80
	_, cha := test(n1, n2)
	fmt.Printf("cha:%v",cha)
}
