package main

import "fmt"

//从键盘读入个数不确定的整数，并判断读入的正数和负数的个数，输入为0时，结束程序
func main() {
	var num, positive, negative int
	for {
		fmt.Printf("输入一个整数：")
		fmt.Scanf("%d", &num)

		if num == 0 {
			break
		} else if num > 0 {
			positive++
			continue
		}
		negative++

	}
	fmt.Printf("共有%d个正数，%d个负数", positive, negative)
}
