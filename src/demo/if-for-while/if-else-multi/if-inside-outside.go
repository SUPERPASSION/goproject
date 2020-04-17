package main

import "fmt"

//跑步，超过8秒进入决赛，否则淘汰；
//根据性别提示进入男子组或者女子组。输入成绩和性别，进行判断。

func main() {
	var speed float64
	var sex string
	fmt.Println("输入你的性别（男或女)：")
	fmt.Scanln(&sex)
	fmt.Println("输入你的成绩：")
	fmt.Scanln(&speed)

	if speed < 8 {
		if sex == "男" {
			fmt.Println("恭喜你，进入男子组决赛")
		} else {
			fmt.Println("恭喜你，进入女子组决赛")
		}
	} else {
		fmt.Println("抱歉，你没有进入决赛")
	}
}
