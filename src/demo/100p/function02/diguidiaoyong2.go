package main

import "fmt"

//已知f(1)=3,f(n)=2*f(n-1)+1;

func tyler01(n int) int {
	if n ==1 {
		return 3
	} else {
		return 2*tyler01(n-1)+1
	}
}

func main() {
	res := tyler01(2)
	fmt.Println(res)
}
