package main

import "fmt"

//func main() {
//	var a int32 = -100
//	var b int32 = 99
//	if a+b >= 50 {
//		fmt.Println("hello world")
//	} else {
//		fmt.Println("hello your head")
//	}
//}

//func main(){
//	var a float64 = 13.14
//	var b float64 = 6.66
//	if a > 10.0 && b < 20.0{
//		fmt.Println(a+b)
//	}
//}

//func main() {
//	var a int32 = 10
//	var b int32 = 20
//	if (a+b)%3 == 0 && (a+b)%5 == 0 {
//		fmt.Println("Pass")
//	}
//}

func main() {
	var a int = 2020
	if (a%4 == 0 && a%100 != 0) || a%400 == 0 {
		fmt.Println("是闰年")
	}
}
