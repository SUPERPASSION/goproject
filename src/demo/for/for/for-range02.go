package main

import "fmt"

func main() {
	str := "hello world!!啦啦"
	for index, val := range str {
		fmt.Printf("index=%d,val=%c\n", index, val)
	}
}
