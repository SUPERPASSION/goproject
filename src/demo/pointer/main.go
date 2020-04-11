package main

import "fmt"

//1、基本数据类型，变量存的就是值，也叫值类型
//
//2、获取变量的地址，用&，比如，var num int，获取num的地址：&num
//
//3、指针类型，变量存的是一个地址，这个地址指向的空间存的才是值
//
//比如：var ptr *int = &num
//
//4、获取指针类型所指向的值，使用：*，比如：var ptr *int，使用*ptr获取p指向的值

func main() {
	//基本数据类型在内存中的布局
	var i int = 10
	fmt.Println(&i)

	//ptr是一个指针变量
	//ptr的类型是*int
	//ptr本身的值是&i
	var ptr *int = &i
	fmt.Println(ptr)
	fmt.Println(*ptr)

}
