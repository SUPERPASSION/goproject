package main

import (
	"fmt"
	"strconv"
)

func main() {
	//var str string = "true"
	//var isB bool
	//isB, _ = strconv.ParseBool(str)
	//fmt.Printf("isB - Type : %T \nisB - Value : %t\n\n", isB, isB)

	var str string = "123"
	var n1 int64
	n1, _ = strconv.ParseInt(str, 10, 64)
	fmt.Printf("n1 - Type : %T \nn1 - Value : %d\n\n", n1, n1)

	var str2 string = "123.456"
	var f1 float64
	f1, _ = strconv.ParseFloat(str2, 64)
	fmt.Printf("f1 - Type : %T \nf1 - Value : %v\n\n", f1, f1)

}
