package main

import "fmt"

func main() {
	//打印该月份所属的季节，3,4,5春季，6,7,8夏季，9,10,11秋季，12,1,2冬季
	var month uint8
	fmt.Printf("输入月份：")
	fmt.Scanln(&month)

	switch month {
	case 3,4,5:
		fmt.Println("春")
	case 6,7,8:
		fmt.Println("夏")
	case 9,10,11:
		fmt.Println("秋")
	case 12,1,2:
		fmt.Println("冬")
	default:
		fmt.Println("Are you fucking on the Earth?")
	}
}
