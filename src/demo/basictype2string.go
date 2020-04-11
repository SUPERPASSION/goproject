package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1 int64 = 100
	num2 := 1.111
	isB := true
	var myChar byte = '1'
	var str string
	fmt.Printf("num1:%T\nnum2:%T\nisB:%T\nmychar:%T\n", num1, num2, isB, myChar)

	str = strconv.FormatInt(num1,10)
	fmt.Printf("str Type : %T,str value : %q\n", str, str)

	str = strconv.FormatFloat(num2,'f',10,64)
	fmt.Printf("str Type : %T,str value : %q\n", str, str)

	str = strconv.FormatBool(isB)
	fmt.Printf("str Type : %T,str value : %q\n", str, str)

	//
	//str = fmt.Sprintf("%d", num1)
	//fmt.Printf("str Type : %T,str value : %q\n", str, str)
	//
	//str = fmt.Sprintf("%f", num2)
	//fmt.Printf("str Type : %T,str value : %q\n", str, str)
	//
	//str = fmt.Sprintf("%t", isB)
	//fmt.Printf("str Type : %T,str value : %q\n", str, str)
	//
	//str = fmt.Sprintf("%c", myChar)
	//fmt.Printf("str Type : %T,str value : %q\n", str, str)
}
