package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i == 8 {
			return
		}
		fmt.Printf("okok-%d\n", i)
	}
}
