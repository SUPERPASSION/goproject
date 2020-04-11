package main

import "fmt"

func main() {
	//单变量声明
	//方法一
	var num1 int
	num1 = 100
	fmt.Printf("num1 is %d\n", num1)

	var num2 int = 200
	fmt.Printf("num1 is %d\n", num2)

	//方法二
	var name1 = "Tyler"
	fmt.Printf("name is %s\n", name1)

	//方法三
	name2 := "Tyler"
	age := 100
	fmt.Printf("%s age is %d\n", name2, age)

	name2 = "Wu"
	fmt.Printf("%s age is %d\n", name2, age)

	//多变量声明
	//方法一
	var name3, name4, name5 string
	name3, name4, name5 = "Tyler", "James", "Allen"

	name3, name4, name5 = "Kobe", "Pau", "Finsher"
	fmt.Println(name3, name4, name5)

	//方法二
	var name6, name7 = "Tyler", "James"
	fmt.Println(name6, name7)

	//方法三
	var (
		num3 int = 100
		num4 int = 200
	)
	fmt.Println(num3, num4)
}
