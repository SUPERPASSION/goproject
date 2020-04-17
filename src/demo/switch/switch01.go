package main

import "fmt"

func main() {
	var week byte
	fmt.Printf("请输入周几（abcdefg）:")
	fmt.Scanf("%c", &week)

	switch week {
	case 'a':
		fmt.Printf("今天周1")
	case 'b':
		fmt.Printf("今天周2")
	case 'c':
		fmt.Printf("今天周3")
	case 'd':
		fmt.Printf("今天周4")
	case 'e':
		fmt.Printf("今天周5")
	case 'f':
		fmt.Printf("今天周6")
	case 'g':
		fmt.Printf("今天周7")
	default:
		fmt.Printf("XJ8输？")
	}
}
