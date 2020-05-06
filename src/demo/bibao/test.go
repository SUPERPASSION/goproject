package main

import (
	"fmt"
	"strings"
)

//编写一个函数makeSuffix(suffix string)，可以接收一个文件后缀名，如.jpg，并返回一个闭包
//调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀，如.jpg，则返回 文件名.jpg，如果有则返回文件名
//要求使用闭包的方式完成
//strings.HasSuffix，该函数可以判断某个字符串是否有指定的后缀

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		//如果name没有指定的后缀，则加上，否则直接返回name
		if strings.HasSuffix(name, suffix) == false {
			return name + suffix
		}
		return name
	}
}

func main() {
	f := makeSuffix(".jpg")
	fmt.Println(f("winter"))  //winter.jpg
	fmt.Println(f("tyler.jpg")) //tyler.jpg
}
