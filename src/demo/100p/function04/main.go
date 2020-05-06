package main

import "fmt"

func cal(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}
func main() {
	sum, sub := cal(10, 5)
	fmt.Println(sum, sub)
}
