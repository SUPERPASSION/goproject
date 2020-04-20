package main

import "fmt"

func main() {
	var max_num int = 9
	for a := 1; a <= max_num; a++ {
		for b := 1; b <= a; b++ {
			fmt.Printf("%v*%v=%v\t", a, b, a*b)
		}
		fmt.Println("")
	}
}
