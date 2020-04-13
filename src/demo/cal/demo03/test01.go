package main

import "fmt"

func main() {
	//假如还有97天放假，问：多少星期多少天
	days_rest_all := 97
	var week_rest_all int
	var days_rest int

	week_rest_all = days_rest_all / 7
	days_rest = days_rest_all % 7

	fmt.Printf("还有%d天放假，即%d星期%d天", days_rest_all, week_rest_all, days_rest)

	//定义一个变量保存华氏温度，华氏温度转换为摄氏温度的公式为：5/9*(华氏温度-100)，求出华氏温度对应的摄氏温度
	var huashi float64 = 200
	var sheshi float64
	//这个地方需要用float类型，因为如果用整数的话5/9得出的结果是0，值就永远等于0
	sheshi = 5.0 / 9 * (huashi - 100)
	fmt.Printf("华氏温度%v，对应的摄氏温度是%v", huashi, sheshi)
}
