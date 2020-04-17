package main

import "fmt"

func main() {
	//%b 表示输出二进制的占位符
	var i int = 255
	fmt.Printf("%b\n", i)

	//直接输出八进制以0开头
	var j int = 011
	fmt.Println(j)

	//直接输出十六进制以0x或0X开头
	var k int = 0x11
	fmt.Println(k)
}
