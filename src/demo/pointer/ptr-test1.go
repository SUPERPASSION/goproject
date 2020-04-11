package main

import "fmt"

func main() {
	var num int = 100
	//fmt.Printf("num - Address:%p\n",&num)

	var ptr *int = &num
	*ptr = 200
	fmt.Printf("num - Address:%v\n", &num)
	fmt.Printf("num - Value:%v\n", num)

	var a int = 100
	var b int = 200
	var ptr2 *int = &a
	*ptr2 = 200
	ptr2 = &b
	*ptr2 = 300
	fmt.Printf("a - Value = %v\nb - Value = %v", a, b)
}
