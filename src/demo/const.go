package main

import "fmt"

func main() {
	const name = "Tyler"
	fmt.Printf("%s\n", name)

	//常量定义之后不允许修改

	const (
		name1 = "Tyler"
		PATH  = "www.tylerwu.top"
		age   = 100
		//一组常量中，如果某个常量没有初始值，就和上一行一致
		w
	)
	fmt.Printf("%s + %s + %d", name1, PATH, w)
}
