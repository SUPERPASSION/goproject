package main

import "fmt"

func main() {
	var age int = 40
	//&&
	//是否大于30，小于50，是，执行
	if age > 30 && age < 50 {
		fmt.Println("OKOK1")
	}
	//是否大于30，小于35，否，不执行
	if age > 30 && age < 35 {
		fmt.Println("OKOK2")
	}

	//||
	//是否大于30，或小于35,，满足第一个，执行
	if age > 30 || age < 35 {
		fmt.Println("OKOK3")
	}

	//!
	//是否 不 大于30，否，不执行
	if !(age > 30) {
		fmt.Println("OKOK4")
	}
	//是否 不 小于30，是，执行
	if !(age < 30) {
		fmt.Println("OKOK5")
	}
}
