package main

import "fmt"

//0+6=6
//1+5=6
//2+4=6
//3+3=6
//...
//6+0=6

//func main() {
//	var a int8
//	var b int8
//	for a = 0; a <= 6; a++ {
//		for b = 0; b <= 6; b++ {
//			if a+b == 6 {
//				fmt.Printf("%v\t+\t%v\t=\t6\n", a, b)
//			}
//		}
//	}
//}

func main() {
	num := 6
	for i := 0; i <= num; i++ {
		fmt.Printf("%v+%v=%v\n", i, num-i, num)
	}
}
