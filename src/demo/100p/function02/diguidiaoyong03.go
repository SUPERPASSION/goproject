package main

import "fmt"

//有一堆桃子，猴子第一天吃其中的一半，并再多吃一个
//以后每天猴子都吃其中的一半，并再多吃一个
//第十天时，想吃了，发现只有一个了，求最初有多少个桃子

func chitaozi(days int) int {
	if days == 10 {
		return 1
	} else{
		return (chitaozi(days+1)+1)*2
	}
}

func main() {
	sum := chitaozi(1)
	fmt.Println(sum)
}
