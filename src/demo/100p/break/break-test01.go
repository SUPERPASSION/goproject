package main

import "fmt"

func main() {
	//100以内的数求和，求出 当和第一次大于20时候的当前数
	//var sum int
	//for i := 1; i < 100; i++ {
	//	sum += i
	//	if sum > 20 {
	//		fmt.Println(i, sum)
	//		break
	//	}
	//}

	//实现登录验证，有三次机会，如果用户名为张无忌，888提示登录成功，否则提示还有几次机会
	var user string
	var password string
	var count int = 0
	for {
		fmt.Printf("请输入用户名：")
		fmt.Scanf("%s", &user)
		fmt.Printf("请输入密码：")
		fmt.Scanf("%s", &password)
		count++
		if user == "张无忌" && password == "888" {
			fmt.Println("登录成功")
		} else if count == 3 {
			fmt.Println("快把手机还给人家")
			break
		} else {
			fmt.Printf("你还剩%d次机会\n", 3-count)
		}
	}

}
