package main

import "fmt"

func main() {
	//使用switch把小写类型的char转为大写（键盘输入），只转换abcde，其他的输出为other
	var char01 byte
	fmt.Printf("请输入一个字符串a/b/c/d/e：")
	fmt.Scanf("%c", &char01)

	switch char01 {
	case 'a':
		fmt.Println("A")
	case 'b':
		fmt.Println("B")
	case 'c':
		fmt.Println("C")
	case 'd':
		fmt.Println("D")
	case 'e':
		fmt.Println("E")
	default:
		fmt.Printf("other")
	}
}
