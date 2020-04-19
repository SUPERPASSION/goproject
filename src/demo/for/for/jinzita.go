package main

import "fmt"

//打印金字塔

func main() {
	var cengshu int = 9
	for a := 1; a <= cengshu; a++ {
		for c := 1; c <= cengshu-a; c++ {
			fmt.Printf(" ")
		}
		for b := 1; b <= 2*a-1; b++ {
			fmt.Print("*")
		}
		fmt.Printf("\n")
	}
}
