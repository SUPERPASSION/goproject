package main

import "fmt"

func main() {
	//输入星期，输出对应值
	var week string
	fmt.Printf("输入星期几：")
	fmt.Scanf("%s", &week)

	switch week {
	case "星期一", "一", "1":
		fmt.Println("干煸豆角")
	case "星期二", "二", "2":
		fmt.Println("醋溜土豆")
	case "星期三", "三", "3":
		fmt.Println("红烧狮子头")
	case "星期四", "四", "4":
		fmt.Println("油炸花生米")
	case "星期五", "五", "5":
		fmt.Println("蒜蓉扇贝")
	case "星期六", "六", "6":
		fmt.Println("东北乱炖")
	case "星期七", "七", "7":
		fmt.Println("大盘鸡")
	}
}
