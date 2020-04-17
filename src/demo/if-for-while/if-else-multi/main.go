package main

import "fmt"

func main() {
	var score int
	fmt.Printf("文捷的成绩是:")
	fmt.Scanln(&score)
	if score == 100 {
		fmt.Println("BMW")
	} else if score > 80 && score < 100 {
		fmt.Println("ip7 plus")
	} else if score >= 60 && score <= 80 {
		fmt.Println("ipad")
	} else {
		fmt.Println("You get nothing , bitch!!")
	}
}
