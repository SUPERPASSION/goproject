package main

import "fmt"

func main() {
	var num int = 100
	//fmt.Printf("num - Address:%p\n",&num)

	var ptr *int = &num
	*ptr = 200
	fmt.Printf("num - Address:%v\n",&num)
	fmt.Printf("num - Value:%v\n",num)
}
