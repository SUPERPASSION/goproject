package main

import "fmt"

func main() {
	//求两个数的最大值
	//a := 200
	//b := 100
	////c := 300
	//if a > b {
	//	fmt.Println(a)
	//}else{
	//	fmt.Println(b)
	//}

	//求三个数的最大值
	a := 2002
	b := 100
	c := 300
	if a > b && a > c {
		fmt.Println("最大是a")
	} else if b > a && b > c {
		fmt.Println("最大是b")
	} else {
		fmt.Println("最大是c")
	}
}