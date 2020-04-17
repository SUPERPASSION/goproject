package main

import "fmt"

func main() {
	//男的高于180cm，财产大于10m，并且帅
	//同时满足3个，嫁
	//满足一个，嫁吧
	//都不满足，不嫁

	var height int
	var money int
	var sex bool
	fmt.Println("请输入你的身高（xxcm），财产，帅不帅（true/false）")
	fmt.Scanln(&height,&money,&sex)
	if height > 180 && money > 1000000 && sex {
		fmt.Println("嫁就完事了")
	} else if height > 180 || money > 1000000 || sex{
		fmt.Println("勉强嫁吧")
	} else {
		fmt.Println("?????")
	}
}
