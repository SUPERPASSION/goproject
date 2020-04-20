package main

import "fmt"

func main() {
label2:
	for i := 0; i <= 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break label2
			}
			fmt.Println("j=", j)
		}
	}
}
