package main

import "fmt"

//打印1~100之间所有9的倍数的整数的个数及总和
func main() {
	var count int
	var sum int

	//此处是踩的坑，for的判断结果必须为true，否则就停止循环
	//for i := 1; i%9 == 0; i++ {
	//	if i <= 100 {
	//		count++
	//		sum += i
	//	}
	//}

	for i := 1; i <= 100; i++ {
		if i%9 == 0 {
			count++
			sum += i
		}
	}
	fmt.Printf("个数:%v", count)
	fmt.Printf("总和:%v", sum)
}
