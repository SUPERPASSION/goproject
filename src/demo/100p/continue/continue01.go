package main

import "fmt"

//用for和continue输出1~100的奇数
func main() {
	for i := 1; i <= 100; i++ {
		if i%2==0{
			continue
		}
		fmt.Printf("%d\n",i)
	}
}
