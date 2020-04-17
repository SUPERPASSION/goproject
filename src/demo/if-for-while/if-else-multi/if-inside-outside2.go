package main

import (
	"fmt"
)

//出票系统，根据淡旺季的月份和年龄，打印票价
//4-10旺季：成人(18-60):60，儿童(<18):半价，老人(>60):1/3
//淡季：成人40，其他半价

func main() {
	var month float32
	var age byte
	fmt.Printf("请输入当前月份（1-12）：")
	fmt.Scanln(&month)
	const high_season = 60
	const low_season = 40
	if month >= 4 && month <= 10 {
		fmt.Printf("请输入您的年龄：")
		fmt.Scanln(&age)
		if age < 18 {
			fmt.Printf("月份：%v\t旺季儿童票价：%v", month, high_season/2)
		} else if age <= 60 {
			fmt.Printf("月份：%v\t旺季成人票价：%v", month, high_season)
		} else {
			fmt.Printf("月份：%v\t旺季老人票价：%v", month, high_season/3)
		}
	} else if (month >= 1 && month < 4) || (month > 10 && month <= 12) {
		fmt.Printf("请输入您的年龄：")
		fmt.Scanln(&age)
		if age >= 18 && age <= 60 {
			fmt.Printf("月份：%v\t淡季成人票价：%v", month, low_season)
		} else {
			fmt.Printf("月份：%v\t淡季其他票价：%v", month, low_season/2)
		}
	} else {
		fmt.Println("你生活在地球吗")
	}
}
