package main

import "fmt"

//func main() {
//	var str string = "hello world"
//
//	for i := 0; i < len(str); i++ {
//		fmt.Printf("%c", str[i])
//	}
//}

func main() {
	var str string = "hello world嗷嗷"
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c\n", str2[i])
	}
}