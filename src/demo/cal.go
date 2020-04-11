package main

import (
	"fmt"
	"unsafe"
)

func main(){
	//var num1 int = 100
	//var num2 int = 99
	//var num_result int = num1 + num2
	//fmt.Printf("%d\n",num_result)

	//var str1 string = "Kobe"
	//var str2 string = "Bryant"
	//var str_result string = str1 + str2
	//fmt.Printf("%s\n",str_result)

	var age int8 = 24
	fmt.Println("age所占用了%d个字节",unsafe.Sizeof(age))
}