package main

import "fmt"

func main() {
	//某人有100，000元，每经过一次路口，需要交费，规则如下
	//当现金>50000时，交费5%
	//当现金<=50000时，交费1000
	//计算该人可以经过多少次路口，使用for break

	var count int
	var money float64 = 100000
	for ; money > 50000; money *= 0.95 {
		count++
	}
	count += int(money / 1000)

	fmt.Printf("介个倒霉蛋可以通过%d次路口", count)
}
