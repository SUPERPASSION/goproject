package main

import "fmt"

//1,1,2,3,5,8,13
//给你一个整数，求出它的斐波那契数

func feibonaqi(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return feibonaqi(n-1) + feibonaqi(n-2)
	}
}

func main() {
	res := feibonaqi(5)
	fmt.Println(res)
}
