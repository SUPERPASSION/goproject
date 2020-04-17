package main

import (
	"fmt"
)

func main() {
	//成绩大于60分，输出合格，低于60，输出不合格，不能大于100
	var score float64
	fmt.Printf("你考几分：")
	fmt.Scanln(&score)

	switch {
	case score >= 60 && score <= 100:
		fmt.Println("合格")
	case score < 60:
		fmt.Println("不合格")
	case score > 100:
		fmt.Println("输入有误")
	}
}
