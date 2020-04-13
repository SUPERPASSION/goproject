package main

import "fmt"

func main() {
	//如果参与运算的数，都是整数，那么除法处理后，去掉小数点的部分，不会四舍五入
	fmt.Println(10 / 4)

	//接收的部分是float类型，也收到的是整数
	var n1 float64 = 10 / 4
	fmt.Println(n1)

	//如果希望保留小数部分，则需要有float类型的值，参与除法运算
	var n2 float64 = 10.0 / 4
	fmt.Println(n2)
}
