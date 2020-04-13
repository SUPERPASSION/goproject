package main

import "fmt"

func main() {
	//演示%的使用
	//%的公式，a % b = a - a / b * b
	n1 := 10 % 3
	n2 := -10 % 3
	n3 := 10 % -3
	n4 := -10 % -3
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
}
