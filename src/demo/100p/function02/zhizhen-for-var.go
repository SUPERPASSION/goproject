package main

import "fmt"

func test01(n1 int) {
	n1 = n1 + 10
}

func test02(n1 *int) {
	*n1 = *n1 + 10
}

func main() {
	num := 20
	test01(num)
	fmt.Printf("%v\n", num)

	num2 := 20
	test02(&num2)
	fmt.Printf("%v\n", num2)
}
