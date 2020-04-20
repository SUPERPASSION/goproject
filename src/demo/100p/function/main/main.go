package main

import "fmt"
import sb "demo/100p/function/cal"

func main() {
	var n1 float64 = 1.1
	var n2 float64 = 5.5
	var operator byte = '/'
	res := sb.Cal(n1,n2,operator)
	fmt.Println("result=",res)
}
