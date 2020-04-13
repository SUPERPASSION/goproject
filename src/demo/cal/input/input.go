package main

import "fmt"

func main() {
	var name string
	var age byte
	var sal float64
	var isPass bool

	fmt.Println("请输入姓名")
	fmt.Scanln(&name)

	fmt.Println("请输入年龄")
	fmt.Scanln(&age)

	fmt.Println("请输入薪水")
	fmt.Scanln(&sal)

	fmt.Println("是否通过考试")
	fmt.Scanln(&isPass)

	fmt.Printf("姓名是%v\n年龄是%v\n薪水是%v\n是否通过考试:%v", name, age, sal, isPass)
}
