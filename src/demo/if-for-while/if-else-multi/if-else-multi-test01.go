package main

import (
	"fmt"
	"math"
)

func main() {
	//求ax²+bx+c=0方程的根，abc分别为函数的参数，如果b²-4ac>0，则有两个解
	//b²-4ac=0，则有一个解；b²-4ac<0，则无解；
	//提示1:x1=(-b+Sqrt(b²-4ac))/2a
	//提示2:x2=(-b-Sqrt(b²-4ac))/2a
	//math.Sqrt(num)；可以求平方根，需要引入math包

	var a, b, c float64 = 3.0, 100.0, 6.0
	m := b*b - 4*a*c

	if m > 0 {
		x1 := (-b + math.Sqrt(m)) / 2 * a
		x2 := (-b - math.Sqrt(m)) / 2 * a
		fmt.Printf("x1=%v\nx2=%v", x1, x2)
	} else if m == 0 {
		x1 := (-b + math.Sqrt(m)) / 2 * a
		fmt.Printf("x1=%v", x1)
	} else {
		fmt.Println("无解")
	}

}
