package main

import "fmt"

func main() {
	var age int
	fmt.Printf("Input your real age:")
	fmt.Scanln(&age)

	if age > 18 {
		fmt.Println("Your age is over 18 you should be responsible for yourself")
	} else{
		fmt.Println("Not over 18 , you can leave")
	}
}