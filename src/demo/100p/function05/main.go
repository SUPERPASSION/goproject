package main

import "fmt"

func sum(n1 int, args ...int) int {
	sum := n1
	//遍历args
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func main() {
	res := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(res)
}
