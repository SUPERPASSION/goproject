package main

import "fmt"

func main() {
	var max, sum, line_count = 100, 0, 0
	for num := 1; num < max; num++ {
		for num2 := 2; num2 <= num; num2++ {
			if num%num2 == 0 && num != num2 {
				break
			} else if num == num2 {
				sum += num
				fmt.Printf("%v\t", num)
				line_count++
				if line_count == 5 {
					fmt.Printf("\n")
					line_count = 0
				}
			}
		}
	}
	fmt.Printf("100以内素数的总和是%v", sum)
}
